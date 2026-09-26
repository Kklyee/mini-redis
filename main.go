package main

import (
	"errors"
	"fmt"
)

func main() {
	store := NewStore()

	_, err := store.Get("hello")

	fmt.Println(err)
	fmt.Println(err == ErrKeyNotFound)
	fmt.Println(errors.Is(err, ErrKeyNotFound))
}
