// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/Zubayear/song-ql/graph"
	"github.com/Zubayear/song-ql/logger"
	"github.com/Zubayear/song-ql/repository"
)

// Injectors from wire.go:

func DependencyProvider() (*graph.Resolver, error) {
	zapLogger := logger.LoggerProvider()
	songRepository, err := repository.DatabaseImlProvider()
	if err != nil {
		return nil, err
	}
	resolver, err := graph.ResolverProvider(zapLogger, songRepository)
	if err != nil {
		return nil, err
	}
	return resolver, nil
}
