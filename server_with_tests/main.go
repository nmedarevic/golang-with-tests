package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

type InMemoryStore struct {
}

func (store *InMemoryStore) GetPlayerScore(name string) int {
	return 123
}

func (store *InMemoryStore) RecordScore(name string) {}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordScore(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		p.SaveScore(w, r)
	case http.MethodGet:
		p.ShowScore(w, r)
	}
}

func (p *PlayerServer) SaveScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	p.store.RecordScore(player)
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) ShowScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Println(player)

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	fmt.Fprint(w, score)
}

func GetPlayerScore(name string) string {
	if name == "Alice" {
		return "20"
	}

	if name == "Bob" {
		return "10"
	}

	return ""
}

func main() {
	server := &PlayerServer{store: &InMemoryStore{}}
	log.Fatal(http.ListenAndServe(":3000", server))
}
