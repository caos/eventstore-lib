package models

type Aggregate interface {
	GetType() string
	GetID() string
	GetVersion() string
	ToFilters() []Filter
}
