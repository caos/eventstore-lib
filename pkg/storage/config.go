package storage

type Config interface {
	New() Storage
}

// Option is an optional extension of the storage layer (e.g. for caching)
type Option func(Storage) error
