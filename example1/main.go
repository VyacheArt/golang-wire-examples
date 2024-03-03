package main

type KeyValueStorage interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

func main() {
	storage := initializeStorage()
	_ = storage.Set("foo", "bar")

	value, _ := storage.Get("foo")
	println(value)

	if value != "bar" {
		println("value is not bar")
	}
}
