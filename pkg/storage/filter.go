package storage

import "github.com/caos/eventstore-lib/pkg/models"

type FilterFunc func(Query)

func BuildFilters(query Query, filters []FilterFunc) {
	for _, f := range filters {
		f(query)
	}
}

func GenerateFilter(operation models.Operation, fieldname string, value interface{}) FilterFunc {
	return func(query Query) {
		query.Condition(fieldname, operation, value)
	}
}
