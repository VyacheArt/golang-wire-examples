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
			wire.Bind(new(KeyValueStorage), new(*inmemory.Storage)),
			inmemory.NewStorage,
		),
	)
}
