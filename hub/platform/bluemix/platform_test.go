package bluemix_test

import (
    "testing"
    "time"

    "github.com/acbodine/iot/hub/event"
    "github.com/acbodine/iot/hub/event/local"
    "github.com/acbodine/iot/hub/platform"
)

const (
    skipMessage = "Environment isn't configured for testing against Bluemix."
)

func TestBluemixPlatformer(t *testing.T) {

    // Get an Bluemix Platformer.
    p, err := platform.Factory(platform.Bluemix)
    if err != nil {
        t.Skip(skipMessage)
    }

    pumped := make(chan *event.Event)

    // Tell Platformer to start pumping Events.
    if err := p.Pump(pumped); err != nil {
        t.Fatal(err)
    }

    // Create a mock Eventer, and register with Platformer.
    e := local.MemoryEventer(time.Millisecond * 10)
    if err := p.Register(e); err != nil {
        t.Fatal(err)
    }

    // TODO: What to do from here? We need to verify that messages will be
    // sent to an MQTT gateway. Can we do this without actually connecting
    // to Bluemix, just for this test?

    if err := p.Terminate(); err != nil {
        t.Fatal(err)
    }
}
