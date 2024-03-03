package main

type KeyValueStorage interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

type Config struct {
	Name string
}

func main() {
	res, clean, err := initializeStorage(Config{Name: "my-name-over-config"})
	if err != nil {
		panic(err)
	}

	println("initial length", res.Length)
	println("initial name", res.Name)

	storage := res.Storage
	_ = storage.Set("foo", "bar")

	value, _ := storage.Get("foo")
	println(value)

	if value != "bar" {
		println("value is not bar")
	}

	clean()
}
