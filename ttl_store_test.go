package main

import (
	"errors"
	"testing"
	"time"
)

func TestTTlStore(t *testing.T) {
	ttlStore := NewTTLStore(time.Second * 5)
	ttlStore.Set("name","alice")

	val, err := ttlStore.Get("name")

	if errors.Is(err,ErrKeyNotFound) {
	  t.Fatal("expected key to exsited")
	}

	if val != "alice" {
	  t.Fatalf("expected value to be alice, got %s", val)
	}



}
