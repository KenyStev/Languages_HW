package main

import (
		"github.com/go-martini/martini"
		// "github.com/martini-contrib/binding"
		"github.com/martini-contrib/render"
		"net/http"
		// "github.com/martini-contrib/cors"
    "log"
    "./services"
	)

const(
  SORTPATH = "resources/mergesort/"
)

func main() {
  m := martini.Classic()

  m.Use(render.Renderer(render.Options{
      Directory:  "templates",
      Extensions: []string{".html"},
  }))

  m.Group("/api", func(mr martini.Router) {
  	// mr.Post("/bitcode/image",CryptImage)

    mr.Get("/sort", func(render render.Render, log *log.Logger) {
      render.HTML(200, "emails_upload", nil)
    })

    mr.Post("/sort", func(writer http.ResponseWriter,r *http.Request) (int, string) {
      log.Println("parsing form")
      err := r.ParseMultipartForm(100000)
      if err != nil {
          return http.StatusInternalServerError, err.Error()
      }

      files := r.MultipartForm.File["files"]
      if val,err := services.Upload(files[0],SORTPATH); err != "ok" {
        return val,err
      }

      sortedFile := services.SortEmails(files[0].Filename)
      log.Println(sortedFile)
      services.Download(sortedFile,files[0].Filename+".sorted",writer)

      return 200, "ok"
    })
  })

  m.Get("/", func() string {
    return "Hello world!"
  })
  m.Get("/hello/:name", func(params martini.Params) string {
	return "Hello " + params["name"]
  })
  m.Run()
}