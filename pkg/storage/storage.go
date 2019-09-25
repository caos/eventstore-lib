package storage

import (
	"context"

	"github.com/caos/eventstore-lib/pkg/models"
)

// Storage is the interface for database interactions
type Storage interface {
	Start(options ...Option) error
	Health() error

	//NEW
	LockAggregates(ctx context.Context, aggregates ...models.Aggregate) error
	ValidateLatestSequence(ctx context.Context, aggregates ...models.Aggregate) error
	PushEvents(ctx context.Context, aggregates ...models.Aggregate) error
	UnlockAggregates(ctx context.Context, aggregates ...models.Aggregate) error
}
