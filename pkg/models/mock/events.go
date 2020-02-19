package mock

import models "github.com/caos/eventstore-lib/pkg/models"

type MockEvents []*MockEvent

func (e *MockEvents) Len() int {
	return len(*e)
}
func (e *MockEvents) Get(index int) models.Event {
	return (*e)[index]
}
func (e *MockEvents) GetAll() []models.Event {
	events := make([]models.Event, len(*e))
	for i, event := range *e {
		events[i] = event
	}
	return events
}
func (e *MockEvents) Append(event models.Event) {
	mockEvent, ok := event.(*MockEvent)
	if !ok {
		return
	}
	*e = append(*e, mockEvent)
	return
}
func (e *MockEvents) Insert(position int, event models.Event) {
	if len(*e) < position {
		e.Append(event)
	}
	mockEvent, ok := event.(*MockEvent)
	if !ok {
		return
	}
	events := (*e)[:position]
	events = append(events, mockEvent)
	events = append(events, (*e)[position+1:]...)

	*e = events
	return
}
func (e *MockEvents) Iterate(f func(event models.Event) error) error {
	for _, event := range *e {
		if err := f(models.Event(event)); err != nil {
			return err
		}
	}
	return nil
}
