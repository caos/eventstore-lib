package eventstore

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/caos/eventstore-lib/pkg/models"
	"github.com/caos/eventstore-lib/pkg/storage"
	caos_errs "github.com/caos/utils/errors"
)

func (es *Service) createStorageFilter(eventType reflect.Type, filters ...models.Filter) ([]storage.FilterFunc, error) {
	filterFuncs := make([]storage.FilterFunc, len(filters))
	for idx, filter := range filters {
		field, found := eventType.FieldByNameFunc(buildFieldnameFilter(filter.GetField()))
		if !found {
			return nil, caos_errs.ThrowInvalidArgument(nil, "EVEN-0VE0d", fmt.Sprintf("fieldname \"%v\" not in event", filter.GetField()))
		}

		filterFunc, err := es.stor.BuildFilter(filter, field)
		if err != nil {
			return nil, err
		}

		filterFuncs[idx] = filterFunc
	}
	return filterFuncs, nil
}

func buildFieldnameFilter(filterFieldname string) func(string) bool {
	return func(fieldname string) bool {
		return strings.ToLower(fieldname) == strings.ToLower(filterFieldname)
	}
}
