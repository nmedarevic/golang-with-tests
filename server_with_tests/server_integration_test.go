package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingScoreAndFetchingThem(t *testing.T) {
	t.Run("Should save and retrieve score", func(t *testing.T) {
		store := NewInMemoryPlayerStore()
		server := PlayerServer{store}
		player := "Bob"

		server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
		server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
		server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))

		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response.Code, http.StatusOK)
	})
}
