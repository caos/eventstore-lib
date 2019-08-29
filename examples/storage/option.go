package storage

import (
	"github.com/caos/eventstore-lib/pkg/storage"
	"github.com/caos/utils/cache"
	caos_err "github.com/caos/utils/errors"
	trace "github.com/caos/utils/tracing"
)

type traceOption interface {
	SetTracer(tracer trace.Tracer)
}

func CreateTracerOption(tracer trace.Tracer) storage.Option {
	return func(storage storage.Storage) error {
		traceStorage, ok := storage.(traceOption)
		if !ok {
			return caos_err.ThrowUnimplementedf(nil, "bTJcn", "SetTracer is unimpelemented in storage %T", storage)
		}
		traceStorage.SetTracer(tracer)
		return nil
	}
}

type cacheOption interface {
	SetCache(cacher cache.Cache)
}

func CreateCacheOption(cacher cache.Cache) storage.Option {
	return func(storage storage.Storage) error {
		cacheStorage, ok := storage.(cacheOption)
		if !ok {
			return caos_err.ThrowUnimplementedf(nil, "4zfcv", "Set Cache is unimplemented in storage %T", storage)
		}
		cacheStorage.SetCache(cacher)
		return nil
	}
}
