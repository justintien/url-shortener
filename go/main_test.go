package main

import (
	"testing"

	"shortener"
)

var a shortener.App

func TestMain(t *testing.T) {
	a := shortener.App{}
	go serve(a)
}
