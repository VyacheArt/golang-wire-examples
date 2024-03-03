//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"wire-tutorial/inmemory"
)

func initializeStorage(cfg Config) KeyValueStorage {
	panic(
		wire.Build(
			wire.FieldsOf(new(Config), "Name"),

			wire.Bind(new(KeyValueStorage), new(*inmemory.Storage)),
			inmemory.NewStorageWithName,
		),
	)
}
