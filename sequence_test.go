package eventstore

import (
	"context"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/caos/eventstore-lib/pkg/models"
	models_mock "github.com/caos/eventstore-lib/pkg/models/mock"
	"github.com/caos/eventstore-lib/pkg/storage"
	storage_mock "github.com/caos/eventstore-lib/pkg/storage/mock"
)

func TestCreateSequenceFilters(t *testing.T) {
	ctrl := gomock.NewController(t)
	events := make([]models.Event, 2)
	for i := 0; i < len(events); i++ {
		event := models_mock.NewMockEvent(ctrl)
		event.EXPECT().GetSequence().Return(uint64(i * 100))

		aggregate := models_mock.NewMockAggregate(ctrl)
		aggregate.EXPECT().GetID().Return("1")
		aggregate.EXPECT().GetType().Return("TestEvent_" + strconv.Itoa(i))
		aggregate.EXPECT().GetVersion().Return("v1.0.0")

		aggregate.EXPECT().ToFilters().Return(nil)
		event.EXPECT().GetAggregate().Return(aggregate)
		events[i] = event
	}
	svc := new(Service)
	svc.setCreateSequenceFilters()
	checks, err := svc.createSequenceFilters(events...)
	require.NoError(t, err)
	_ = checks
	// storage := storage_mock.NewMockStorage(ctrl)

}

func TestIsLatestSequences(t *testing.T) {
	ctrl := gomock.NewController(t)
	svc := new(Service)
	svc.setIsLatestSequences()
	store := storage_mock.NewMockStorage(ctrl)
	svc.stor = store

	filters := make([]lastSequenceCheck, 1)
	for i := uint64(0); int(i) < len(filters); i++ {
		Filters := []storage.Filter{func(storage.Query) {}}
		filters[i] = lastSequenceCheck{eventSequence: i, filters: Filters}
		store.EXPECT().IsLatestSequence(context.Background(), gomock.Any(), gomock.Any()).DoAndReturn(func(context.Context, interface{}, interface{}) bool {
			<-time.After(time.Duration(rand.Intn(60)) * time.Millisecond)
			return true
		})
	}

	isLatest := svc.isLatestSequences(context.Background(), filters...)
	assert.True(t, isLatest)
}

func TestIsLatestSequencesNotLatest(t *testing.T) {
	ctrl := gomock.NewController(t)
	svc := new(Service)
	svc.setIsLatestSequences()
	store := storage_mock.NewMockStorage(ctrl)
	svc.stor = store

	filters := make([]lastSequenceCheck, 1)
	for i := uint64(0); int(i) < len(filters); i++ {
		Filters := []storage.Filter{func(storage.Query) {}}
		filters[i] = lastSequenceCheck{eventSequence: i, filters: Filters}
		store.EXPECT().IsLatestSequence(context.Background(), gomock.Any(), gomock.Any()).DoAndReturn(func(context.Context, interface{}, interface{}) bool {
			<-time.After(time.Duration(rand.Intn(60)) * time.Millisecond)
			return false
		})
	}

	isLatest := svc.isLatestSequences(context.Background(), filters...)
	assert.False(t, isLatest)
}
