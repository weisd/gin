package main

import (
	"github.com/weisd/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {

		data := ctx.Input().Query("da")

		ctx.String(200, data)
	})

	r.Run(":8000")
}
