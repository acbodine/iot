package local

import (
    "sync"
    "time"

    "github.com/acbodine/iot/hub/event"
)

type mock struct {
    events      chan *event.Event
    terminate   chan bool

    timeout     time.Duration

    sync.WaitGroup
}

func MockEventer(timeout time.Duration) event.Eventer {
    if timeout == 0 {
        timeout = time.Millisecond * 10
    }

    return &mock{
        terminate:  make(chan bool),

        timeout:    timeout,
    }
}

func (m *mock) Emit(events chan *event.Event) error {
    m.events = events

    m.Add(1)
    go func () {
        defer m.Done()

        var timer *time.Timer = time.NewTimer(m.timeout)

        for {
            select {

            case <- m.terminate:
                return

            case <- timer.C:
                m.events <- event.New([]byte(""), time.Now())
            }

            timer.Reset(m.timeout)
        }
    }()

    return nil
}

func (m *mock) Terminate() error {
    defer close(m.terminate)
    m.terminate <- true

    m.Wait()

    m.events = nil

    return nil
}
