package storage

import "github.com/caos/eventstore-lib/pkg/models"

type Filter func(Query)

func BuildFilters(query Query, filters []Filter) {
	for _, f := range filters {
		f(query)
	}
}

func GenerateFilter(operation models.Operation, fieldname string, value interface{}) Filter {
	return func(query Query) {
		query.Condition(fieldname, operation, value)
	}
}
