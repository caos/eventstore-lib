package storage

//go:generate mockgen -package mock -destination ./mock/storage.mock.go github.com/caos/eventstore-lib/pkg/storage Storage
//go:generate mockgen -package mock -destination ./mock/query.mock.go github.com/caos/eventstore-lib/pkg/storage Query
