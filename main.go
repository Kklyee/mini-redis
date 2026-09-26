package main

func main() {
	store := NewStore()

	// 1. Set name = alice
	store.Set("name", "alice")

	// 2. Get name
	// 打印 value 和 ok
	val, ok := store.Get("name")
	println(val, ok)

	// 3. Delete name
	store.Delete("name")

	val, ok = store.Get("name")
	println(val, ok)
}
