package bluemix

import (
    "sync"

    "github.com/acbodine/iot/hub/event"
)

type platform struct {
    events      chan *event.Event

    client      *client

    sync.WaitGroup
}

// Provide platform.Factory() an easy constructor.
func Platform() (*platform, error) {
    c, err := Connect(nil)
    if err != nil {
        return nil, err
    }

    p := &platform{
        events:     make(chan *event.Event),
        client:     c,
    }

    return p, nil
}

// Implement platform.Platformer interface.
func (p *platform) Register(e event.Eventer) error {
    if err := e.Emit(p.events); err != nil {
        return err
    }

    return nil
}

// Implement platform.Platformer interface.
func (p *platform) Unregister(e event.Eventer) error {
    if err := e.Terminate(); err != nil {
        return err
    }

    return nil
}

// Implement the platform.Platformer interface.
func (p *platform) Pump(ledger chan *event.Event) error {
    return nil
}

// Implement the platform.Platformer interface.
func (p *platform) Terminate() error {
    close(p.events)

    // Wait for any goroutines that are pumping Events.
    p.Wait()

    p.client.Disconnect()

    return nil
}
