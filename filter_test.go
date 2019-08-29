package eventstore

import (
	"reflect"
	"testing"

	"github.com/caos/utils/errors"

	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/caos/eventstore-lib/pkg/models"
	models_mock "github.com/caos/eventstore-lib/pkg/models/mock"
	"github.com/caos/eventstore-lib/pkg/storage"
	storage_mock "github.com/caos/eventstore-lib/pkg/storage/mock"
)

func TestBuildFieldnameFilter(t *testing.T) {
	filter := buildFieldnameFilter("aSdFgH")
	assert.True(t, filter("AsDfGh"))
	assert.False(t, filter("jkl"))
	assert.False(t, filter("asdfg"))
}

func TestCreateStorageFilter(t *testing.T) {
	ctrl := gomock.NewController(t)
	svc := new(Service)
	stor := storage_mock.NewMockStorage(ctrl)
	svc.stor = stor

	var filters []models.Filter
	typeFilter := models_mock.NewMockFilter(ctrl)
	typeFilter.EXPECT().GetField().Return("type")
	commandFilter := models_mock.NewMockFilter(ctrl)
	commandFilter.EXPECT().GetField().Return("command")
	filters = append(filters, typeFilter, commandFilter)

	eventType := reflect.TypeOf(testEvent{})
	typeField, ok := eventType.FieldByName("Type")
	require.True(t, ok)
	commandField, ok := eventType.FieldByName("Command")
	require.True(t, ok)
	stor.EXPECT().BuildFilter(typeFilter, typeField).Return(func(storage.Query) {}, nil)
	stor.EXPECT().BuildFilter(commandFilter, commandField).Return(func(storage.Query) {}, nil)

	Filters, err := svc.createStorageFilter(eventType, filters...)
	require.NoError(t, err)
	assert.Len(t, Filters, 2)
}

func TestCreateStorageFilterWrongFieldname(t *testing.T) {
	ctrl := gomock.NewController(t)
	svc := new(Service)
	stor := storage_mock.NewMockStorage(ctrl)
	svc.stor = stor

	var filters []models.Filter
	wrongFilter := models_mock.NewMockFilter(ctrl)
	wrongFilter.EXPECT().GetField().Return("wrong").Times(2)
	filters = append(filters, wrongFilter)

	eventType := reflect.TypeOf(testEvent{})
	Filters, err := svc.createStorageFilter(eventType, filters...)

	require.Error(t, err)
	assert.True(t, errors.IsErrorInvalidArgument(err))
	assert.Empty(t, Filters)
}
