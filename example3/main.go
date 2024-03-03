package main

type KeyValueStorage interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Name() string
}

type Config struct {
	Name string
}

func main() {
	storage := initializeStorage(Config{Name: "name-over-config"})
	println(storage.Name())

	_ = storage.Set("foo", "bar")

	value, _ := storage.Get("foo")
	println(value)

	if value != "bar" {
		panic("value is not bar")
	}
}
