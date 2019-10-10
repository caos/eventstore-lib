package models

type Filters interface {
	Limit() uint64
	Filters() []Filter
}

type Filter interface {
	GetField() FilterEventKey
	GetOperation() Operation
	GetValue() interface{}
}

type FilterEventKey int32

type Operation int32
