package storage

import (
	"context"
	"reflect"

	"github.com/caos/eventstore-lib/pkg/models"
)

// Storage is the interface for database interactions
type Storage interface {
	Start(options ...Option) error
	// GetEvent searches an event by the given eventID
	GetEvent(ctx context.Context, event models.Event, eventID string) error
	// IsLatestSequence checks if the given sequence is the lastest by the filters.
	// "filters" contains all aggregate filters
	IsLatestSequence(ctx context.Context, sequence uint64, filters ...FilterFunc) bool
	// GetEvents searches for a list of events by the given filters
	GetEvents(ctx context.Context, events models.Events, filters ...FilterFunc) error
	// CreateEvents stores the given events.
	CreateEvents(ctx context.Context, events ...models.Event) error
	// Health checks the availability of the storage
	Health() error
	// BuildFilter converts the given filter to a storage filter function
	// structField is for mapping the event-field to the database field
	BuildFilter(filter models.Filter, structField reflect.StructField) (FilterFunc, error)
}