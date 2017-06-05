package event

import (
    "time"
)

// To keep things simple for now, Event is purposefully abstract
// and contains a generic []byte payload.
type Event struct {
    payload     []byte
    timestamp   time.Time
}

func New(p []byte, t time.Time) *Event {
    return &Event{
        payload:    p,
        timestamp:  t,
    }
}

type Eventer interface {

    // Emit provides a starting function for the underlying
    // Event aggregation logic, and the Eventer should start
    // send Events on the provided *Event channel.
    Emit(chan *Event) error

    // Terminate forces an Eventer to stop sending Events.
    Terminate() error

}
