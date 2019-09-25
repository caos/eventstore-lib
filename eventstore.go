package eventstore

import (
	"context"

	"github.com/caos/eventstore-lib/pkg/models"
	"github.com/caos/eventstore-lib/pkg/storage"
)

type Eventstore interface {
	Start() error
	Health() error

	PushEvents(ctx context.Context, events ...models.Aggregate) error
	// GetEvents(ctx context.Context, events models.Events, filters models.EventFilter) error
}

var _ Eventstore = (*Service)(nil)

type Service struct {
	store storage.Storage
}

func (es *Service) Start() error {
	return es.store.Start()
}

func (es *Service) Health() error {
	return es.store.Health()
}
