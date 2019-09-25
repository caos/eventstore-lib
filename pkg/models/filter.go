package models

type Filters interface {
	AggregateType() string
	AggregateID() string
	Filters() []Filter
}

type Filter interface {
	GetField() string
	GetOperation() Operation
	GetValue() interface{}
}
