package eventstore

import (
	"context"
	"reflect"

	"github.com/caos/eventstore-lib/pkg/models"
	"github.com/caos/eventstore-lib/pkg/storage"
	caos_errs "github.com/caos/utils/errors"
	"github.com/caos/utils/tracing"
)

func (es *Service) CreateEvents(ctx context.Context, events ...models.Event) (err error) {
	ctx, span := tracing.NewServerSpan(ctx)
	defer func() { span.EndWithError(err) }()

	if events == nil {
		return nil
	}

	sequenceFilters, err := es.createSequenceFilters(events...)
	if err != nil {
		return err
	}

	if isLatest := es.isLatestSequences(ctx, sequenceFilters...); !isLatest {
		return caos_errs.ThrowAlreadyExists(nil, "EVENT-RXWTQ", "sequence wrong")
	}

	return es.stor.CreateEvents(ctx, events...)
}

func (es *Service) GetEvent(ctx context.Context, event models.Event, eventID string) (err error) {
	ctx, span := tracing.NewServerSpan(ctx)
	defer func() { span.EndWithError(err) }()

	return es.stor.GetEvent(ctx, event, eventID)
}

func (es *Service) GetEvents(ctx context.Context, events models.Events, eventFilters models.EventFilter) (err error) {
	ctx, span := tracing.NewServerSpan(ctx)
	defer func() { span.EndWithError(err) }()

	eventType := reflect.TypeOf(events).Elem().Elem().Elem() // *[]*Event
	filters := make([]storage.Filter, 0)

	storageFilters, err := es.createStorageFilter(eventType, eventFilters.GetFilters()...)
	if err != nil {
		return err
	}
	filters = append(filters, storageFilters...)

	aggregateFilters, err := es.createStorageFilter(eventType, eventFilters.GetAggregate().ToFilters()...)
	if err != nil {
		return err
	}
	filters = append(filters, aggregateFilters...)

	return es.stor.GetEvents(ctx, events, filters...)
}
