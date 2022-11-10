package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/hi", nil)
	router().ServeHTTP(w, req)

	expected := "Go la la!"
	actual := w.Body.String()
	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}
