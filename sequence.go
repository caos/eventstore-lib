package eventstore

import (
	"context"
	"reflect"
	"sync"

	"github.com/caos/eventstore-lib/pkg/models"
	"github.com/caos/eventstore-lib/pkg/storage"
)

type lastSequenceCheck struct {
	eventSequence uint64
	filters       []storage.Filter
}

func (svc *Service) setCreateSequenceFilters() {
	svc.createSequenceFilters = func(events ...models.Event) (check []lastSequenceCheck, err error) {
		check = make([]lastSequenceCheck, len(events))
		for i, event := range events {
			check[i].eventSequence = event.GetSequence()
			check[i].filters, err = svc.createStorageFilter(reflect.TypeOf(event).Elem(), event.GetAggregate().ToFilters()...)
			if err != nil {
				return nil, err
			}
		}
		return check, nil
	}
}

func (svc *Service) setIsLatestSequences() {
	svc.isLatestSequences = func(ctx context.Context, checks ...lastSequenceCheck) bool {
		isLatest := make(chan bool)
		go func() {
			var wg sync.WaitGroup
			for _, check := range checks {
				wg.Add(1)
				go func(sequenceCheck lastSequenceCheck) {
					defer wg.Done()
					isLatest <- svc.stor.IsLatestSequence(ctx, sequenceCheck.eventSequence, sequenceCheck.filters...)
				}(check)
			}
			wg.Wait()
			close(isLatest)
		}()
		for latest := range isLatest {
			if !latest {
				return false
			}
		}
		return true
	}
}
