package inmemory

type Storage struct {
	name string
	data map[string]string
}

func NewStorage() *Storage {
	return NewStorageWithName("default")
}

func NewStorageWithName(name string) *Storage {
	return &Storage{
		name: name,
		data: make(map[string]string, 10),
	}
}

func NewStorageWithNameAndClean(name string) (*Storage, func()) {
	s := NewStorageWithName(name)
	return s, func() {
		s.data = nil
		println("storage", name, "cleaned")
	}
}

func (s *Storage) Get(key string) (string, error) {
	return s.data[key], nil
}

func (s *Storage) Set(key string, value string) error {
	s.data[key] = value
	return nil
}

func (s *Storage) Name() string {
	return s.name
}

func (s *Storage) Len() int {
	return len(s.data)
}
