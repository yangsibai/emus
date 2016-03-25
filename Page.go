package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Page struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Title     string        `json:"title" bson:"title"`
	URL       string        `json:"URL" bson:"URL"`
	Content   string        `json:"content" bson:"content"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	Deleted   bool          `json:"_" bson: "deleted"`
}
