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
	err = C.Find(bson.M{"deleted": bson.M{"$ne": true}}).Sort("-created_at").All(&pages)
	return
}

func DeletePage(id string) error {
	session := GetSession()
	defer session.Close()
	C := session.DB("emu").C("page")
	change := bson.M{"$set": bson.M{"deleted": true}}
	err := C.Update(bson.M{"_id": bson.ObjectIdHex(id)}, change)
	return err
}
