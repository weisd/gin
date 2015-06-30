package gin

import (
	"net/http"
	"net/url"
)

func (c *Context) GetCookie(name string) string {
	cookie, err := c.Request.Cookie(name)
	if err != nil {
		return ""
	}
	val, _ := url.QueryUnescape(cookie.Value)
	return val
}

func (c *Context) SetCookie(name string, value string, others ...interface{}) {
	cookie := &http.Cookie{}
	cookie.Name = name
	cookie.Value = url.QueryEscape(value)

	if len(others) > 0 {
		switch v := others[0].(type) {
		case int:
			cookie.MaxAge = v
		case int64:
			cookie.MaxAge = int(v)
		case int32:
			cookie.MaxAge = int(v)
		}
	}

	cookie.Path = "/"
	if len(others) > 1 {
		if v, ok := others[1].(string); ok && len(v) > 0 {
			cookie.Path = v
		}
	}

	if len(others) > 2 {
		if v, ok := others[2].(string); ok && len(v) > 0 {
			cookie.Domain = v
		}
	}

	if len(others) > 3 {
		switch v := others[3].(type) {
		case bool:
			cookie.Secure = v
		default:
			if others[3] != nil {
				cookie.Secure = true
			}
		}
	}

	if len(others) > 4 {
		if v, ok := others[4].(bool); ok && v {
			cookie.HttpOnly = true
		}
	}

	// http.SetCookie(c.Writer, cookie)
	// c.Request.AddCookie(cookie)

	c.Writer.Header().Add("Set-Cookie", cookie.String())
}
