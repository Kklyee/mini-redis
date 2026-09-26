package main

import "time"

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

func (s *TTLStore) Get(key string)(TTLEntry,bool) {
   val, ok := s.data[key]

   if !ok {
     return TTLEntry{}, ok
   }
   return val,ok
}
