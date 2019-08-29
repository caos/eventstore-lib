package eventstore

import (
	"context"

	"github.com/caos/eventstore-lib/pkg/storage"
	"github.com/caos/utils/logging"
	tracing "github.com/caos/utils/tracing/config"
)

type Config struct {
	Store  storage.Config
	Tracer *tracing.TracingConfig
}

func (c *Config) New() Eventstore {
	es := new(Service)
	es.setCreateSequenceFilters()
	es.setIsLatestSequences()

	if c.Store == nil {
		logging.Log("EVENT-PLdu7").Panic("no storage config provided")
	}
	es.stor = c.Store.NewStorage()
	if c.Tracer != nil && c.Tracer.Config != nil {
		err := c.Tracer.Config.NewTracer(context.TODO())
		logging.Log("EVENT-Cqqlx").OnError(err).Error("unable to create tracer")
	} else {
		logging.Log("EVENT-L5Y6T").Info("no tracing config provided")
	}

	return es
}
