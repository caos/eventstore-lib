package models

import "time"

type Aggregate interface {
	GetType() string
	GetID() string
	GetVersion() string
	ToFilters() []Filter
}

type Agg interface {
	Type() string
	Version() string
	ID() string
	Events() []Eve
	LatestSequence() uint64
	LatestEventDate() time.Time
}

type Eve interface {
	ID() string

	CreationDate() time.Time
	SetCreationDate(time.Time)

	Command() string

	// Sequence() uint64
	SetSequence(sequence uint64)

	GetData() []byte
}
