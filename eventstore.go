package eventstore

import (
	"context"

	"github.com/caos/eventstore-lib/pkg/models"
	"github.com/caos/eventstore-lib/pkg/storage"
)

type Eventstore interface {
	Health() error

	CreateEvents(ctx context.Context, events ...models.Event) error
	GetEvents(ctx context.Context, events models.Events, filters models.EventFilter) error
	GetEvent(ctx context.Context, event models.Event, id string) error
}

var _ Eventstore = (*Service)(nil)

type Service struct {
	stor                  storage.Storage
	createSequenceFilters func(events ...models.Event) (check []lastSequenceCheck, err error)
	isLatestSequences     func(ctx context.Context, checks ...lastSequenceCheck) bool
}

func (es *Service) Health() error {
	return es.stor.Health()
}
