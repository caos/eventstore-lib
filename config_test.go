package eventstore

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/caos/eventstore-lib/pkg/storage"
)

func TestNew(t *testing.T) {
	assert.Panics(t, newEventstorePanic())

	config := new(Config)
	config.Store = new(storageConfig)

	es := config.New()
	assert.NotNil(t, es)
}

type storageConfig struct{}

func (c *storageConfig) NewStorage() storage.Storage {
	return nil
}

func newEventstorePanic() func() {
	return func() {
		conf := new(Config)
		conf.New()
	}
}
