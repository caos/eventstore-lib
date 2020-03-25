package eventstore

import (
	"context"

	"github.com/caos/eventstore-lib/pkg/models"
	"github.com/caos/eventstore-lib/pkg/repository"
)

type Eventstore interface {
	Health(ctx context.Context) error

	PushEvents(ctx context.Context, events ...models.Aggregate) error
	Filter(ctx context.Context, events models.Events, query models.SearchQuery) error
}

var _ Eventstore = (*Service)(nil)

type Service struct {
	repo repository.Repository
}

func (es *Service) Health(ctx context.Context) error {
	return es.store.Health(ctx)
}
