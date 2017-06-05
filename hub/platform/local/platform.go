package local

import (
    "sync"

    "github.com/acbodine/iot/hub/event"
)

type platform struct {
    events      chan *event.Event

    sync.WaitGroup
}

// Provide platform.Factory() an easy constructor.
func Platform() (*platform, error) {
    p := &platform{
        events:     make(chan *event.Event),
    }

    return p, nil
}

// Implement the platform.Platformer interface.
func (p *platform) Register(e event.Eventer) error {
    if err := e.Emit(p.events); err != nil {
        return err
    }

    return nil
}

// Implement the platform.Platformer interface.
func (p *platform) Unregister(e event.Eventer) error {
    if err := e.Terminate(); err != nil {
        return err
    }

    return nil
}

// Implement the platform.Platformer interface.
func (p *platform) Pump(ledger chan *event.Event) error {
    p.Add(1)

    go func () {
        defer p.Done()

        for {
            select {
            case e, ok := <- p.events:

                // If we received an Event, put it in the ledger.
                if e != nil && ledger != nil {
                    ledger <- e
                }

                // If the channel closed, then we are done pumping.
                if !ok {
                    return
                }
            }
        }
    }()

    return nil
}

func (p *platform) Terminate() error {
    close(p.events)

    // Wait for any goroutines that are pumping Events.
    p.Wait()

    return nil
}
