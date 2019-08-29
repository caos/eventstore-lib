package eventstore

import (
	"context"

	"github.com/caos/eventstore-lib/pkg/models"
	"github.com/caos/eventstore-lib/pkg/storage"
)

type testEvent struct {
	Type    string
	Command string
}

func newTestService(stor storage.Storage) Eventstore {
	svc := new(Service)
	svc.stor = stor
	svc.isLatestSequences = newTestIsLatestSequences()
	svc.createSequenceFilters = newTestCreateSequenceFilters()

	return svc
}

func newTestIsLatestSequences() func(ctx context.Context, checks ...lastSequenceCheck) bool {
	return nil
}

func newTestCreateSequenceFilters() func(events ...models.Event) (check []lastSequenceCheck, err error) {
	return nil
}
