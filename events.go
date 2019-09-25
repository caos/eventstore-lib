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

	for _, aggregate := range aggregates {
		for idx, event := range aggregate.Events().GetAll() {
			event.SetSequence(aggregate.LatestSequence() + uint64(idx) + 1)
		}
	}

	return es.store.PushEvents(ctx, aggregates...)
}

// func (es *Service) GetEvents(ctx context.Context, events models.Events, eventFilters models.EventFilter) (err error) {
// 	ctx, span := tracing.NewServerSpan(ctx)
// 	defer func() { span.EndWithError(err) }()

// 	eventType := reflect.TypeOf(events).Elem().Elem().Elem() // *[]*Event
// 	filters := make([]storage.Filter, 0)

// 	storageFilters, err := es.createStorageFilter(eventType, eventFilters.GetFilters()...)
// 	if err != nil {
// 		return err
// 	}
// 	filters = append(filters, storageFilters...)

// 	aggregateFilters, err := es.createStorageFilter(eventType, eventFilters.GetAggregate().ToFilters()...)
// 	if err != nil {
// 		return err
// 	}
// 	filters = append(filters, aggregateFilters...)

// 	return es.store.GetEvents(ctx, events, filters...)
// }
