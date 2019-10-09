package models

type Filters interface {
	Limit() uint64
	Filters() []Filter
}

type Filter interface {
	GetField() uint64
	GetOperation() Operation
	GetValue() interface{}
}
