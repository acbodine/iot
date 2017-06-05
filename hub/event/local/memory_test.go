package local_test

import (
    "sync"
    "testing"
    "time"

    "github.com/acbodine/iot/hub/event"
    "github.com/acbodine/iot/hub/event/local"
)

func TestMemoryEventer(t *testing.T) {
    var wg sync.WaitGroup

    events := make(chan *event.Event)

    m := local.MemoryEventer(time.Millisecond * 1)

    wg.Add(1)
    go func () {
        defer wg.Done()

        returned := []*event.Event{}

        for len(returned) < 1 {
            returned = append(returned, <- events)
        }
    }()

    if err := m.Emit(events); err != nil {
        t.Fatal(err)
    }

    // Guarantee that we received an *event.Event.
    wg.Wait()

    if err := m.Terminate(); err != nil {
        t.Fatal(err)
    }
}
