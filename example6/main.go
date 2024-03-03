package main

type KeyValueStorage interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

type Config struct {
	Name string
}

func main() {
	res, clean := initializeStorage(Config{Name: "name-over-config"})
	storage := res.Storage
	println(res.Name)
	println(res.Capacity)

	_ = storage.Set("foo", "bar")

	value, _ := storage.Get("foo")
	println(value)

	if value != "bar" {
		panic("value is not bar")
	}

	clean()
}
