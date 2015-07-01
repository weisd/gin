// weisd
package gzip

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/weisd/gin"
)

const (
	HeaderAcceptEncoding  = "Accept-Encoding"
	HeaderContentEncoding = "Content-Encoding"
	HeaderContentLength   = "Content-Length"
	HeaderContentType     = "Content-Type"
	HeaderVary            = "Vary"
)

// Gziper returns a Handler that adds gzip compression to all requests.
// Make sure to include the Gzip middleware above other middleware
// that alter the response body (like the render middleware).
func Gziper() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// fmt.Println("gzip")
		if !strings.Contains(ctx.Request.Header.Get(HeaderAcceptEncoding), "gzip") {
			return
		}

		headers := ctx.Writer.Header()
		headers.Set(HeaderContentEncoding, "gzip")
		headers.Set(HeaderVary, HeaderAcceptEncoding)

		gz := gzip.NewWriter(ctx.Writer)
		defer gz.Close()

		gzw := gzipResponseWriter{gz, ctx.Writer}
		ctx.Writer = gzw

		ctx.Next()

		// 防止被当作gzip下载
		if ctx.Writer.Size() == -1 {
			ctx.Writer.Header().Set(HeaderContentType, "text/plain")
			ctx.Writer.Header().Set(HeaderVary, "")
		}

		ctx.Writer.Header().Del("Content-Length")
	}
}

type gzipResponseWriter struct {
	w *gzip.Writer
	gin.ResponseWriter
}

func (grw gzipResponseWriter) Write(p []byte) (int, error) {

	if len(grw.Header().Get(HeaderContentType)) == 0 {
		grw.Header().Set(HeaderContentType, http.DetectContentType(p))
	}

	return grw.w.Write(p)
}

func (grw gzipResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker, ok := grw.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, fmt.Errorf("the ResponseWriter doesn't support the Hijacker interface")
	}
	return hijacker.Hijack()
}
