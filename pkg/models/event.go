package models

import "time"

// Events is an abstraction for a slice of events.
type Events interface {
	Len() int
	Get(index int) Event
	GetAll() []Event
	Append(events Event) Events
	Insert(position int, event Event) Events
	Iterate(func(event Event) error) error
}

// Event represents the minimal representation of an eventstore event
type Event interface {
	GetID() string
	GetCreationDate() time.Time
	SetCreationDate(time.Time)
	GetCommand() string
	SetSequence(sequence uint64)
	GetSequence() uint64
	GetLastSequence() uint64
	SetLastSequence(uint64)
	GetAggregate() Aggregate
	GetData() []byte
}
