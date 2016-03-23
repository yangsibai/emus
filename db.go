package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func GetSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial(config.MongoURL)

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}

func StorePage(page *Page) (err error) {
	session := GetSession()
	defer session.Close()
	C := session.DB("emu").C("page")
	err = C.Insert(&page)
	return
}

func GetPage(id string) (page Page, err error) {
	session := GetSession()
	defer session.Close()
	C := session.DB("emu").C("page")

	oid := bson.ObjectIdHex(id)
	err = C.FindId(oid).One(&page)
	return
}

func GetAllPages() (pages []Page, err error) {
	session := GetSession()
	defer session.Close()
	C := session.DB("emu").C("page")
	err = C.Find(nil).All(&pages)
	return
}
