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

func initializeStorage(cfg Config) (*storageSet, func()) {
	panic(
		wire.Build(
			wire.FieldsOf(new(Config), "Name"),

			wire.Bind(new(KeyValueStorage), new(*inmemory.Storage)),
			inmemory.NewStorageWithNameAndClean,

			provideCapacity,
			provideName,

			wire.Struct(new(storageSet), "*"),
		),
	)
}

func provideCapacity(ims *inmemory.Storage) (int, func()) {
	return ims.Len(), func() {
		println("capacity provider cleaner")
	}
}

func provideName(ims *inmemory.Storage) StorageName {
	return StorageName(ims.Name())
}
