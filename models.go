package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	Link struct {
		Id      bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Title   string        `json:"title" bson:"title"`
		Url     string        `json:"url" bson:"url"`
		Created time.Time     `json:"created" bson:"created"`
		Tags    Tags          `json:"tags" bson:"tags"`
	}
	Links []Link
	Tag   struct {
		Name string `json:"name" bson:"name,omitempty"`
	}
	Tags []Tag
)
