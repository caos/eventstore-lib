package eventstore

import (
	"context"

	"github.com/caos/eventstore-lib/pkg/models"
	"github.com/caos/utils/errors"
	"github.com/caos/utils/tracing"
)

type validtor struct {
	models.Aggregate
}

func (v *validtor) EventCount() int {
	return v.Events().Len()
}

func (es *Service) PushEvents(ctx context.Context, aggregates ...models.Aggregate) (err error) {
	ctx, span := tracing.NewServerSpan(ctx)
	defer func() { span.EndWithError(err) }()

	if aggregates == nil {
		return nil
	}

	for _, aggregate := range aggregates {
		err := es.store.ValidateAndReserveSequence(ctx, &validtor{Aggregate: aggregate})
		if err != nil {
			return errors.ThrowInvalidArgument(err, "EVENT-uwJjT", "sequence wrong")
		}
		for idx, event := range aggregate.Events().GetAll() {
			event.SetSequence(aggregate.LatestSequence() + 1 + uint64(idx))
		}
	}

	return es.store.PushEvents(ctx, aggregates...)
}

/*

Agg: User
ID: "402853029482093"
LS: 0
Events: [
	username
	password
	firstname
]

Agg: Unique_username
ID "rs380"
LS: 0
Events: [
	used
]

1. Validate and Reserve Sequences per aggregate (agg.ID, agg.Type, agg.LatestSequence)
2. Set reserved Sequences on the events
3. write events

*/

// func (es *Service) GetEvent(ctx context.Context, event models.Event, eventID string) (err error) {
// 	ctx, span := tracing.NewServerSpan(ctx)
// 	defer func() { span.EndWithError(err) }()

// 	return es.store.GetEvent(ctx, event, eventID)
// }

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
