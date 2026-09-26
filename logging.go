package main

import (
	"errors"
	"log"
)

type LoggingStorage struct {
	inner  Storage
	logger *log.Logger
}

func NewLoggingStorage(
	inner Storage,
	logger *log.Logger,
) *LoggingStorage {
	return &LoggingStorage{
		inner:  inner,
		logger: logger,
	}
}

func (l *LoggingStorage) Set(key, value string) {
	l.logger.Printf("set %q = %q", key, value)
	l.inner.Set(key, value)
}
func (l *LoggingStorage) Get(key string) (string, error) {
   val,err := l.inner.Get(key)

   if errors.Is(err,ErrKeyNotFound) {
     l.logger.Printf("Get key %q miss",key)
     return "",err
   }

   l.logger.Printf("Get key %q hit", key)
   return val, nil
}
func (l *LoggingStorage) Delete(key string) {
  l.logger.Printf("delete %q", key)
  l.inner.Delete(key)
}
