package main

type Storage interface {
	Get(key string) (string, error)
	Set(key, value string)
	Delete(key string)
}
