package main

import "errors"

var ErrKeyNotFound = errors.New("Key not found")

type Store struct {
	data map[string]string
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Get(key string) (string, error) {
	val, ok := s.data[key]
	if ok {
		return val, nil
	}
	return "", ErrKeyNotFound
}

func (s *Store) Set(key, value string) {
	s.data[key] = value
}

func (s *Store) Delete(key string) {
	delete(s.data, key)
}
