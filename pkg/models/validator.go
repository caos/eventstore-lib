package models

type Validator interface {
	Type() string
	ID() string
	LatestSequence() uint64
	EventCount() int
}
