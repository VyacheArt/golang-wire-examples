//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"wire-tutorial/inmemory"
)

type resultSet struct {
	Storage KeyValueStorage
	Length  int
	Name    StorageName
}

type StorageName string

// possible signatures:
// obj
// obj, error
// obj, cleaner
// obj, cleaner, error

func initializeStorage(cfg Config) (*resultSet, func(), error) {
	panic(
		wire.Build(
			wire.FieldsOf(new(Config), "Name"),

			wire.Bind(new(KeyValueStorage), new(*inmemory.Storage)),
			inmemory.NewStorageWithName,

			provideLength,
			provideName,

			wire.Struct(new(resultSet), "*"),
		),
	)
}

func provideLength(ims *inmemory.Storage) (int, func(), error) {
	return ims.Len(), func() {
		println("provideLength cleaned")
	}, nil
}

func provideName(ims *inmemory.Storage) (StorageName, func(), error) {
	return StorageName(ims.Name()), func() {
		println("provideName cleaned")
	}, nil
}
