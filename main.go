package main

import (
	"net/http"
	gee_v0_2 "write-gin/gee_v0.2"
)

func main() {
	r := gee_v0_2.New()
	r.GET("/index", func(c *gee_v0_2.Context) {
		c.String(http.StatusOK, "你好啊")
	})
	r.GET("/hello", func(c *gee_v0_2.Context) {
		c.Json(200, gee_v0_2.H{
			"hello": "world",
		})
	})
	r.GET("/", func(c *gee_v0_2.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.Run(":8090")

}
