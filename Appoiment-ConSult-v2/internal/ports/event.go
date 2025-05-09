package ports

import "Enrollment/internal/application/events"

type Event interface {
	Publish(eventName string, payload any)
	Subscribe(eventName string, handler events.EventHandler)
}