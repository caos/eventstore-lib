package eventstore

import (
	"github.com/caos/eventstore-lib/pkg/repository"
)

type Config struct {
	Repository repository.Repository
}

func Start(conf Config) Eventstore {
	return &Service{repo: conf.Repository}
}
