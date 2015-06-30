package main

import (
	"github.com/weisd/gin"
	"github.com/weisd/gin/middleware/session"
	_ "github.com/weisd/gin/middleware/session/redis"
)

func main() {
	r := gin.Default()
	r.Use(session.Sessioner(session.Options{Provider: "file"}))

	r.GET("/", func(ctx *gin.Context) {
		sess := session.GetSession(ctx)
		sess.Get("haha")
		ctx.JSON(200, gin.H{"haha": sess.Get("haha")})
	})

	r.GET("/set", func(ctx *gin.Context) {
		sess := session.GetSession(ctx)
		sess.Set("haha", "weisd")
		ctx.JSON(200, gin.H{"haha": "ok"})
	})

	r.Run(":8000")
}
