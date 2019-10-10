package models

type Filters interface {
	Limit() uint64
	Filters() []Filter
}

type Filter interface {
	GetField() Field
	GetOperation() Operation
	GetValue() interface{}
}

type Field int32

type Operation int32
