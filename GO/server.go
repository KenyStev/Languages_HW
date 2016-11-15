package main

import (
		"github.com/go-martini/martini"
		// "github.com/martini-contrib/binding"
		// "github.com/martini-contrib/render"
		// "net/http"
		// "github.com/martini-contrib/cors"
	)

func main() {
  m := martini.Classic()


  m.Group("/api", func(r martini.Router) {
  	r.Post("/bitcode/image",CryptImage)
  })

  m.Get("/", func() string {
    return "Hello world!"
  })
  m.Get("/hello/:name", func(params martini.Params) string {
	return "Hello " + params["name"]
  })
  m.Run()
}