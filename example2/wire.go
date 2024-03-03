//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"wire-tutorial/inmemory"
)

func initializeStorage() KeyValueStorage {
	panic(
		wire.Build(
			wire.Value("new-name"),
			wire.Bind(new(KeyValueStorage), new(*inmemory.Storage)),
			inmemory.NewStorageWithName,
		),
	)
}
