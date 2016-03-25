package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/yangsibai/webutils"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"net/url"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to emus!")
}

func HandleAddPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var page Page
	err := decoder.Decode(&page)
	if err != nil {
		webutils.WriteErrorResponse(w, err)
		return
	}
	page.ID = bson.NewObjectId()
	page.CreatedAt = time.Now()
	err = StorePage(&page)
	if err != nil {
		webutils.WriteErrorResponse(w, err)
		return
	}
	webutils.WriteResponse(w, page.ID.Hex())
}

type PageItem struct {
	Page Page
	Host string
}

// list all pages
func ListPages(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	pages, err := GetAllPages()
	if err != nil {
		panic(err)
	}
	var newPages []PageItem
	for i := 0; i < len(pages); i++ {
		page := pages[i]
		u, err := url.Parse(page.URL)
		if err != nil {
			panic(err)
		}
		item := PageItem{pages[i], u.Host}
		newPages = append(newPages, item)
	}
	ren.HTML(w, http.StatusOK, "pages", newPages)
}

// render a single page by filename
func RenderPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	page, err := GetPage(id)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, page.Content)
}

func HandleDeletePage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	err := DeletePage(id)
	if err != nil {
		webutils.WriteErrorResponse(w, err)
	}
	webutils.WriteResponse(w, id)
}
