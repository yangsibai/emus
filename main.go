package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var ren *render.Render

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to emu")
}

func HandlePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t := time.Now()
	filename := filepath.Join("pages", t.Format("20060102150405")+".html")

	f, err := os.Create(filename)
	check(err)

	defer f.Close()

	f.WriteString(r.FormValue("Content"))
	f.Sync()
	fmt.Fprint(w, "roger that")
}

// list all pages
func ListPages(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	files, err := ioutil.ReadDir("pages")
	if err != nil {
		log.Fatal(err)
	}

	names := make([]string, len(files))

	for _, f := range files {
		names = append(names, f.Name())
	}
	ren.HTML(w, http.StatusOK, "pages", names)
}

// render a single page by filename
func RenderPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ren.HTML(w, http.StatusOK, "page", ps.ByName("name"))
}

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// init: create directory: pages
func init() {
	existed, err := exists("pages")
	check(err)
	if !existed {
		err := os.Mkdir("pages", 0755)
		check(err)
	}
}

func main() {
	ren = render.New(render.Options{
		Directory: "tmpls",
	})

	router := httprouter.New()

	router.GET("/", Index)
	router.POST("/page", HandlePage)
	router.GET("/pages", ListPages)
	router.GET("/page/:name", RenderPage)

	log.Fatal(http.ListenAndServe(":8080", router))
}
