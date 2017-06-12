package local_test

import (
    "sync"
    "testing"
    "time"

    "github.com/acbodine/iot/hub/event"
    "github.com/acbodine/iot/hub/event/local"
    "github.com/acbodine/iot/hub/platform"
)

func TestLocalPlatformer(t *testing.T) {

    // Get a local Platformer.
    p, err := platform.Factory(platform.Local)
    if err != nil {
        t.Fatal(err)
    }

    pumped := make(chan *event.Event)

    // Tell Platformer to start pumping Events.
    if err := p.Pump(pumped); err != nil {
        t.Fatal(err)
    }

    // Create a mock Eventer, and register with Platformer.
    e := local.MockEventer(time.Millisecond * 10)
    if err := p.Register(e); err != nil {
        t.Fatal(err)
    }

    var wg sync.WaitGroup

    wg.Add(1)
    go func () {
        defer wg.Done()

        <- pumped
    }()

    wg.Wait()

    if err := p.Terminate(); err != nil {
        t.Fatal(err)
    }
}
