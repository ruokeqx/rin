package main

import (
	"net/http"
	"rin"
)

// func day1() {
// 	r := rin.New()
// 	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
// 		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
// 	})
// 	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
// 		for k, v := range req.Header {
// 			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
// 		}
// 	})
// 	r.Run(":9999")
// }

func main() {
	r := rin.New()
	r.GET("/", func(c *rin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello rin</h1>")
	})
	r.GET("/hello", func(c *rin.Context) {
		// expect /hello?name=ruokeqx
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.POST("/login", func(c *rin.Context) {
		c.JSON(http.StatusOK, rin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":9999")
}
