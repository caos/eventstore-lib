package eventstore

import (
	"github.com/caos/eventstore-lib/pkg/storage"
	"github.com/caos/utils/logging"
)

type Config struct {
	Store storage.Config
}

func (c *Config) New() Eventstore {
	es := new(Service)
	// es.setCreateSequenceFilters()
	// es.setIsLatestSequences()

	if c.Store == nil {
		logging.Log("EVENT-PLdu7").Panic("no storage config provided")
	}
	es.store = c.Store.New()

	return es
}
