package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
	"github.com/yangsibai/webutils"
	"log"
	"net/http"
)

var config struct {
	Address  string `json:"address"`
	MongoURL string `json:"mongoURL"`
}

var ren *render.Render

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	ren = render.New(render.Options{
		Directory: "tmpls",
	})

	router := httprouter.New()

	router.GET("/", HandleHomePage)
	router.POST("/page", HandleAddPage)
	router.GET("/page/:id", HandleRenderPage)
	router.GET("/meta/:id", HandlePageMeta)
	router.POST("/page/delete/:id", HandleDeletePage)
	router.ServeFiles("/public/*filepath", http.Dir("public"))

	log.Printf("emus is running at %s", config.Address)
	log.Fatal(http.ListenAndServe(config.Address, router))
}

func init() {
	err := webutils.ReadConfig("config.json", &config)
	if err != nil {
		panic(err)
	}
}
