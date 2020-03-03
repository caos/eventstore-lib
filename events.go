package eventstore

import (
	"context"

	"github.com/caos/eventstore-lib/pkg/models"
	"github.com/caos/utils/logging"
	"github.com/caos/utils/tracing"
)

func (es *Service) PushEvents(ctx context.Context, aggregates ...models.Aggregate) (err error) {
	ctx, span := tracing.NewServerSpan(ctx)
	defer func() { span.EndWithError(err) }()

	if aggregates == nil {
		return nil
	}

	if err = es.store.LockAggregates(ctx, aggregates...); err != nil {
		return err
	}

	defer func(lockedAggregates ...models.Aggregate) {
		err := es.store.UnlockAggregates(ctx, lockedAggregates...)
		logging.Log("EVENT-879yy").OnError(err).Error("unable to unlock aggregates")
	}(aggregates...)

	if err = es.store.ValidateLatestSequence(ctx, aggregates...); err != nil {
		return err
	}

	return es.store.PushEvents(ctx, aggregates...)
}

func (es *Service) Filter(ctx context.Context, events models.Events, query models.SearchQuery) (err error) {
	ctx, span := tracing.NewServerSpan(ctx)
	defer func() { span.EndWithError(err) }()

	return es.store.Filter(ctx, events, query)
}
