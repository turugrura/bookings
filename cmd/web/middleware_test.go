package main

import (
	"net/http"
	"testing"
)

func TestWriteToConsole(t *testing.T) {
	var myH myHandler

	h := writeToConsole(&myH)

	switch v := h.(type) {
	case http.Handler:
		//
	default:
		t.Errorf("type is not http.Handler, got %t", v)
	}
}

func TestNoSurf(t *testing.T) {
	var myH myHandler

	h := noSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		//
	default:
		t.Errorf("type is not http.Handler, got %t", v)
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler

	h := sessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
		//
	default:
		t.Errorf("type is not http.Handler, got %t", v)
	}
}
