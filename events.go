package eventstore

import (
	"context"

	"github.com/caos/eventstore-lib/pkg/models"
	"github.com/caos/logging"
)

func (es *Service) PushEvents(ctx context.Context, aggregates ...models.Aggregate) (err error) {
	if aggregates == nil {
		return nil
	}

	if err = es.repo.LockAggregates(ctx, aggregates...); err != nil {
		return err
	}

	defer func(lockedAggregates ...models.Aggregate) {
		err := es.repo.UnlockAggregates(ctx, lockedAggregates...)
		logging.Log("EVENT-879yy").OnError(err).Error("unable to unlock aggregates")
	}(aggregates...)

	if err = es.repo.ValidateLatestSequence(ctx, aggregates...); err != nil {
		return err
	}

	return es.repo.PushEvents(ctx, aggregates...)
}

func (es *Service) Filter(ctx context.Context, events models.Events, query models.SearchQuery) (err error) {
	return es.repo.Filter(ctx, events, query)
}
