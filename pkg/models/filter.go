package models

type EventFilter interface {
	GetAggregate() Aggregate
	GetFilters() []Filter
}

type Filter interface {
	GetField() string
	GetOperation() Operation
	GetValue() interface{}
}
