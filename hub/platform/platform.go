package platform

import (
    "github.com/acbodine/iot/hub/event"
    "github.com/acbodine/iot/hub/platform/bluemix"
    "github.com/acbodine/iot/hub/platform/local"
)

type Platformer interface {

    // Register allows event.Eventers to obtain the *event.Event chan
    // a Platformer is consuming *event.Events from to send to it's
    // corresponding IoT platform.
    Register(e event.Eventer) error

    // Unregister allows business logic to force a Platformer to stop
    // consuming/shipping Events from a specific Eventer.
    // Unregister allows previously registered event.Eventers to notify
    // a Platformer that it is done sending event.Events.
    Unregister(e event.Eventer) error

    // Pump is a blocking call that assures event.Events get sent
    // to the designated IoT platform.
    Pump(ledger chan *event.Event) error

    // Terminate stops the Platformer from pumping event.Events.
    Terminate() error

}

func Factory(t Type) (Platformer, error) {
    switch t {
    case Bluemix:
        return bluemix.Platform()
    default:
        return local.Platform()
    }
}
