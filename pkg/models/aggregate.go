package models

type Aggregate interface {
	Type() string
	ID() string
	Events() Events
	LatestSequence() uint64
}
