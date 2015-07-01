package main

import (
	"fmt"
	"github.com/weisd/gin"
	"github.com/weisd/gin/middleware/gzip"
	"github.com/weisd/gin/middleware/session"
	// _ "github.com/weisd/gin/middleware/session/redis"
)

func main() {
	r := gin.New()
	r.Use(gzip.Gziper())
	r.Use(session.Sessioner())

	r.GET("/", func(ctx *gin.Context) {
		// sess := session.GetSession(ctx)
		// sess.Get("haha")
		// ctx.JSON(200, gin.H{"haha": sess.Get("haha")})

		flash := session.GetFlashValues(ctx)

		ctx.JSON(200, gin.H{"haha": flash.InfoMsg})
	})

	r.GET("/set", func(ctx *gin.Context) {
		// sess := session.GetSession(ctx)
		// sess.Set("haha", "weisd")
		// ctx.JSON(200, gin.H{"haha": "ok"})

		flash := session.GetFlash(ctx)

		flash.Info("weisd")

		fmt.Println("55")
	})

	r.Run(":8000")
}
