package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

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
		p.saveScore(w, r)
	case http.MethodGet:
		p.showScore(w, r)
	}
}

func (p *PlayerServer) saveScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	p.store.RecordScore(player)
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
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
	store := NewInMemoryPlayerStore()
	server := &PlayerServer{store: store}
	log.Fatal(http.ListenAndServe(":3000", server))
}
