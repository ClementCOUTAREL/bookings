package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var mh myHandler

	h := NoSurf(&mh)

	switch v := h.(type) {
	case http.Handler:
		return
	default:
		t.Errorf("type %s is not http.Handler", v)
	}
}

func TestSessionLoad(t *testing.T) {
	var mh myHandler

	h := SessionLoad(&mh)

	switch v := h.(type) {
	case http.Handler:
		return
	default:
		t.Errorf("type %s is not http.Handler", v)
	}
}
