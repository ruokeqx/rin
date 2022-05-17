package main

import (
	"log"
	"net/http"
	"rin"
	"time"
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

// func day2() {
// 	r := rin.New()
// 	r.GET("/", func(c *rin.Context) {
// 		c.HTML(http.StatusOK, "<h1>Hello rin</h1>")
// 	})
// 	r.GET("/hello", func(c *rin.Context) {
// 		// expect /hello?name=ruokeqx
// 		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
// 	})
// 	r.POST("/login", func(c *rin.Context) {
// 		c.JSON(http.StatusOK, rin.H{
// 			"username": c.PostForm("username"),
// 			"password": c.PostForm("password"),
// 		})
// 	})
// 	r.Run(":9999")
// }

// func day3() {
// 	r := rin.New()
// 	r.GET("/", func(c *rin.Context) {
// 		c.HTML(http.StatusOK, "<h1>Hello rin</h1>")
// 	})
// 	r.GET("/hello", func(c *rin.Context) {
// 		// expect /hello?name=ruokeqx
// 		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
// 	})
// 	r.GET("/hello/:name", func(c *rin.Context) {
// 		// expect /hello/ruokeqx
// 		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
// 	})
// 	r.GET("/assets/*filepath", func(c *rin.Context) {
// 		c.JSON(http.StatusOK, rin.H{"filepath": c.Param("filepath")})
// 	})
// 	r.Run(":9999")
// }

// func day4() {
// 	r := rin.New()
// 	r.GET("/index", func(c *rin.Context) {
// 		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
// 	})
// 	v1 := r.Group("/v1")
// 	{
// 		v1.GET("/", func(c *rin.Context) {
// 			c.HTML(http.StatusOK, "<h1>Hello rin</h1>")
// 		})
// 		v1.GET("/hello", func(c *rin.Context) {
// 			// expect /hello?name=ruokeqx
// 			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
// 		})
// 	}
// 	v2 := r.Group("/v2")
// 	{
// 		v2.GET("/hello/:name", func(c *rin.Context) {
// 			// expect /hello/ruokeqx
// 			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
// 		})
// 		v2.POST("/login", func(c *rin.Context) {
// 			c.JSON(http.StatusOK, rin.H{
// 				"username": c.PostForm("username"),
// 				"password": c.PostForm("password"),
// 			})
// 		})
// 	}
// 	r.Run(":9999")
// }

func main() {
	r := rin.New()
	r.Use(rin.Logger()) // global midlleware
	r.GET("/", func(c *rin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello rin</h1>")
	})
	v2 := r.Group("/v2")
	v2.Use(func(c *rin.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *rin.Context) {
			// would not execute,cause Fail Abort
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
