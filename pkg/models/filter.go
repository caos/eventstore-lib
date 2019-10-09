package models

type Filters interface {
	Limit() uint64
	Filters() []Filter
}

type Filter interface {
	GetField() int32
	GetOperation() Operation
	GetValue() interface{}
}
