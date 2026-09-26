package main

import (
	"testing"
	"time"
)

func TestTTlStore(t *testing.T) {
	ttlStore := NewTTLStore(time.Second * 5)
	ttlStore.Set("name","alice")

	val, ok := ttlStore.Get("name")

	if !ok {
	  t.Fatal("expected key to exsited")
	}

	if val.Value != "alice" {
	  t.Fatalf("expected value to be alice, got %s", val.Value)
	}

	if !val.ExpiresAt.After(time.Now()) {
	  t.Fatalf("expire time should be after now")
	}

}
