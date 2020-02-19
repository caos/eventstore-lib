package models

type SearchQuery interface {
	Limit() uint64
	OrderDesc() bool
	Filters() []Filter
}

type Filter interface {
	GetField() Field
	GetOperation() Operation
	GetValue() interface{}
}

type Field int32

type Operation int32
