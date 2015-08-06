package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"time"
)

type (
	LinkController struct {
		s *mgo.Session
	}
)

func NewLinkController(s *mgo.Session) *LinkController {
	return &LinkController{s}
}

func (lc LinkController) GetLinks(w http.ResponseWriter, r *http.Request) {
	links := Links{}
	var query bson.M = bson.M{}
	if urlQuery := r.URL.Query(); len(urlQuery) != 0 {
		query = bson.M{"tags.name": urlQuery.Get("tag")}
	}
	if err := lc.s.DB("links").C("links").Find(query).Sort("-created").All(&links); err != nil {
		log.Fatal(err)
	}
	if err := json.NewEncoder(w).Encode(links); err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
}

func (lc LinkController) CreateLink(w http.ResponseWriter, r *http.Request) {
	link := Link{}
	link.Id = bson.NewObjectId()
	link.Created = time.Now()
	if err := json.NewDecoder(r.Body).Decode(&link); err != nil {
		log.Fatal(err)
	}
	if err := lc.s.DB("links").C("links").Insert(link); err != nil {
		log.Fatal(err)
	}
	lj, _ := json.Marshal(link)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", lj)
}

func (lc LinkController) RemoveLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if !bson.IsObjectIdHex(id) {
		http.NotFound(w, r)
	}
	if err := lc.s.DB("links").C("links").RemoveId(bson.ObjectIdHex(id)); err != nil {
		http.NotFound(w, r)
	}
	w.WriteHeader(http.StatusNoContent)
}
