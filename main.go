package main

import (
	"net/http"
	gee_v0_3 "write-gin/gee_v0.3"
)

func main() {
	r := gee_v0_3.New()
	r.GET("/", func(c *gee_v0_3.Context) {
		c.HTML(http.StatusOK, "<h1>Hello gee_v0_3</h1>")
	})

	r.GET("/hello", func(c *gee_v0_3.Context) {

		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *gee_v0_3.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gee_v0_3.Context) {
		c.JSON(http.StatusOK, gee_v0_3.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
