package main

import (
	"fmt"
	"time"
)

type TTLEntry struct {
	Value     string
	ExpiresAt time.Time
}

type TTLStore struct {
	data map[string]TTLEntry
	ttl  time.Duration
}

func NewTTLStore(ttl time.Duration) *TTLStore {
	return &TTLStore{
		data: make(map[string]TTLEntry),
		ttl:  ttl,
	}
}

func (s *TTLStore) Set(key, value string) {
	s.data[key] = TTLEntry{
		Value:     value,
		ExpiresAt: time.Now().Add(s.ttl),
	}
}

func (s *TTLStore) Get(key string)(string,error) {
   val, ok := s.data[key]

   if !ok {
     return "", fmt.Errorf("get %q: %w",key,ErrKeyNotFound)
   }
   if time.Now().After(val.ExpiresAt) {
     delete(s.data,key)
     return "", fmt.Errorf("get %q: %w",key,ErrKeyNotFound)
   }
   return val.Value,nil
}
