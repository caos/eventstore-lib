package models

// Events is an abstraction for a slice of events.
type Events interface {
	Len() int
	Get(index int) Event
	GetAll() []Event
	Append(events Event)
	Insert(position int, event Event)
}

// Event represents the minimal representation of an eventstore event
type Event interface{}
