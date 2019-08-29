package storage

import "github.com/caos/eventstore-lib/pkg/models"

type Query interface {
	Condition(fieldname string, operation models.Operation, value interface{})
}
