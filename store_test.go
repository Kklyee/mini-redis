package main

import (
	"errors"
	"testing"
)

func TestSetAndGet(t *testing.T) {
	store := NewStore()

	store.Set("name", "alice")

	got, err := store.Get("name")

	if err != nil {
		t.Fatal("expected key to exist")
	}

	if got != "alice" {
		t.Fatalf("expected %q, got %q", "alice", got)
	}
}

func TestGetMissingKey(t *testing.T) {
	store := NewStore()

	_, err := store.Get("missing")

	if !errors.Is(err, ErrKeyNotFound) {
		t.Fatalf("expected ErrKeyNotFound, got %v", err)
	}

}

func TestSetOverwrite(t *testing.T) {
	store := NewStore()
	store.Set("name", "alice")
	store.Set("name", "bob")

	got, err := store.Get("name")

	if errors.Is(err, ErrKeyNotFound) {
		t.Fatal("expected key to exist")
	}

	if got != "bob" {
		t.Fatalf("expected %q, got %q", "bob", got)
	}
}

func TestDelete(t *testing.T) {
	store := NewStore()
	store.Set("name", "alice")
	store.Delete("name")
	val, err := store.Get("name")
	if err == nil {
		t.Fatalf("expected key to be deleted, got %q", val)
	}
}
