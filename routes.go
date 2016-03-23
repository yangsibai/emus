package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/yangsibai/webutils"
	"gopkg.in/mgo.v2/bson"
	"net/http"
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

// list all pages
func ListPages(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	pages, err := GetAllPages()
	if err != nil {
		panic(err)
	}
	ren.HTML(w, http.StatusOK, "pages", pages)
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