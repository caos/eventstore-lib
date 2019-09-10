package eventstore

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/caos/eventstore-lib/pkg/models"
	models_mock "github.com/caos/eventstore-lib/pkg/models/mock"
	"github.com/caos/eventstore-lib/pkg/storage"
	storage_mock "github.com/caos/eventstore-lib/pkg/storage/mock"
	"github.com/caos/utils/errors"
)

func initSvc(t *testing.T) *Service {
	svc := new(Service)
	svc.store = storage_mock.NewMockStorage(gomock.NewController(t))

	return svc
}

func TestGetEvent(t *testing.T) {
	ctx := context.Background()
	event := models_mock.NewMockEvent(gomock.NewController(t))
	eventID := "testID"

	svc := new(Service)
	stor := storage_mock.NewMockStorage(gomock.NewController(t))
	stor.EXPECT().GetEvent(context.Background(), event, eventID).Return(nil)
	svc.store = stor

	err := svc.GetEvent(ctx, event, eventID)
	assert.NoError(t, err)
}

func TestGetEventWrongID(t *testing.T) {
	ctx := context.Background()
	event := models_mock.NewMockEvent(gomock.NewController(t))
	eventID := "i do not exist"

	svc := new(Service)
	stor := storage_mock.NewMockStorage(gomock.NewController(t))
	stor.EXPECT().GetEvent(context.Background(), event, eventID).Return(errors.ThrowNotFound(nil, "EVENT-czr9y", "id not found"))
	svc.store = stor

	err := svc.GetEvent(ctx, event, eventID)
	assert.Error(t, err)
	assert.True(t, errors.IsNotFound(err))
}

func TestCreateEventsNoEvents(t *testing.T) {
	svc := new(Service)
	err := svc.CreateEvents(context.Background())
	assert.NoError(t, err)
}

func TestCreateEventsOneEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	svc := new(Service)
	svc.createSequenceFilters = func(events ...models.Event) (check []lastSequenceCheck, err error) {
		return []lastSequenceCheck{
			lastSequenceCheck{eventSequence: 32, filters: []storage.Filter{func(storage.Query) {}}},
			lastSequenceCheck{eventSequence: 5436, filters: []storage.Filter{func(storage.Query) {}}},
		}, nil
	}
	svc.isLatestSequences = func(ctx context.Context, checks ...lastSequenceCheck) bool {
		<-time.After(time.Duration(rand.Intn(50)) * time.Millisecond)
		return true
	}
	event := models_mock.NewMockEvent(ctrl)
	store := storage_mock.NewMockStorage(ctrl)
	store.EXPECT().CreateEvents(context.Background(), event).Return(nil)
	svc.store = store

	err := svc.CreateEvents(context.Background(), event)
	assert.NoError(t, err)
}

func TestCreateEventsWrongSequence(t *testing.T) {
	svc := new(Service)
	filterCount := 5
	ctrl := gomock.NewController(t)
	svc.createSequenceFilters = func(events ...models.Event) (check []lastSequenceCheck, err error) {
		filters := make([]lastSequenceCheck, filterCount)
		for i := uint64(0); i < uint64(filterCount); i++ {
			filters[i] = lastSequenceCheck{eventSequence: i, filters: []storage.Filter{func(storage.Query) {}}}
		}
		return filters, nil
	}
	svc.isLatestSequences = func(ctx context.Context, checks ...lastSequenceCheck) bool {
		<-time.After(time.Duration(rand.Intn(50)) * time.Millisecond)
		return false
	}

	err := svc.CreateEvents(context.Background(), models_mock.NewMockEvent(ctrl))
	assert.Error(t, err)
	assert.True(t, errors.IsErrorAlreadyExists(err), "should be already exists error")
}

func TestGetEvents(t *testing.T) {
	svc := new(Service)
	ctrl := gomock.NewController(t)

	var events models.Events
	events = new(models_mock.MockEvents)

	aggregate := models_mock.NewMockAggregate(ctrl)
	aggregate.EXPECT().ToFilters().Return(nil)

	var eventFilter models.EventFilter
	filters := make([]models.Filter, 0)
	filter := models_mock.NewMockEventFilter(ctrl)
	filter.EXPECT().GetFilters().Return(filters)
	filter.EXPECT().GetAggregate().Return(aggregate)
	eventFilter = filter

	store := storage_mock.NewMockStorage(ctrl)
	store.EXPECT().GetEvents(context.Background(), events, []storage.Filter{})
	svc.store = store

	err := svc.GetEvents(context.Background(), events, eventFilter)
	assert.NoError(t, err)
}
