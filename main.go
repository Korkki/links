package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"log"
)

func main() {
	r := mux.NewRouter()
	lc := NewLinkController(getSession())
	r.HandleFunc("/links", lc.GetLinks).Methods("GET")
	r.HandleFunc("/links", lc.CreateLink).Methods("POST")
	r.HandleFunc("/links/{id}", lc.RemoveLink).Methods("DELETE")
	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":8000")
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
	return s.Clone()
}
