package models

//go:generate mockgen -package mock -destination ./mock/event.mock.go github.com/caos/eventstore-lib/pkg/models Event
//go:generate mockgen -package mock -destination ./mock/aggregate.mock.go github.com/caos/eventstore-lib/pkg/models Aggregate
//go:generate mockgen -package mock -destination ./mock/filter.mock.go github.com/caos/eventstore-lib/pkg/models Filter,EventFilter
