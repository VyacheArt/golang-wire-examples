//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"wire-tutorial/inmemory"
)

type storageSet struct {
	Storage  KeyValueStorage
	Capacity int
	Name     StorageName
}

type StorageName string

func initializeStorage(cfg Config) *storageSet {
	panic(
		wire.Build(
			wire.FieldsOf(new(Config), "Name"),

			wire.Bind(new(KeyValueStorage), new(*inmemory.Storage)),
			inmemory.NewStorageWithName,

			provideCapacity,
			provideName,

			wire.Struct(new(storageSet), "*"),
		),
	)
}

func provideCapacity(ims *inmemory.Storage) int {
	return ims.Len()
}

func provideName(ims *inmemory.Storage) StorageName {
	return StorageName(ims.Name())
}
