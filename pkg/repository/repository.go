package repository

import (
	"context"

	"github.com/caos/eventstore-lib/pkg/models"
)

// Repository is the interface for database interactions
type Repository interface {
	Health(ctx context.Context) error

	// LockAggregates locks the given aggregates for writes. If this is not needed return nil
	LockAggregates(ctx context.Context, aggregates ...models.Aggregate) error
	// ValidateLatestSequence validate that the given sequence is the latest sequence of the given aggregates.
	// It returns an error if the sequence of ANY given aggregate is out dated
	ValidateLatestSequence(ctx context.Context, aggregates ...models.Aggregate) error
	// PushEvents adds all events of the given aggregates to the eventstreams of the aggregates.
	// This call is transaction save. The transaction will be rolled back if one event fails
	PushEvents(ctx context.Context, aggregates ...models.Aggregate) error
	// UnlockAggregates unlocks the given aggregates for other writes. If this is not needed return nil
	UnlockAggregates(ctx context.Context, aggregates ...models.Aggregate) error
	// Filter returns all events matching the given search query
	Filter(ctx context.Context, events models.Events, query models.SearchQuery) error
}
