package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type PlayerStubStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *PlayerStubStore) GetPlayerScore(name string) int {
	score := s.scores[name]

	return score
}

func (s *PlayerStubStore) RecordScore(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {
	store := &PlayerStubStore{
		scores: map[string]int{
			"Bob":   10,
			"Alice": 20,
		},
	}
	server := &PlayerServer{
		store: store,
	}

	t.Run("returns Bob's score", func(t *testing.T) {
		request := newGetScoreRequest("Bob")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns Alice's score", func(t *testing.T) {
		request := newGetScoreRequest("Alice")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns 404 for missing players", func(t *testing.T) {
		request := newGetScoreRequest("Marko")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestSaveScores(t *testing.T) {
	store := PlayerStubStore{
		scores:   map[string]int{},
		winCalls: nil,
	}
	server := &PlayerServer{&store}
	t.Run("it returns accept on POST", func(t *testing.T) {
		player := "Marko"
		request := newPostScoreRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to record scores, want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("Did not store the correct score. Got %q, want %q", store.winCalls[0], player)
		}
	})
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newPostScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("did not get the correct status! got %d, want %d", got, want)
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
