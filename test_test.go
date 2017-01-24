package main

import (
	"io/ioutil"
	"log"
	"testing"
)

func init() {
	log.SetOutput(ioutil.Discard)
}

func TestSuccess(t *testing.T) {
	if false {
		t.Fatalf("TestSuccess failed.")
	}
}

func TestFailure(t *testing.T) {
	if false {
		t.Fatalf("TestFailure failed.")
	}
}
