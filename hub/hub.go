package main

import (
    "sync"
    "time"

    "github.com/acbodine/iot/hub/event"
    "github.com/acbodine/iot/hub/event/local"
    "github.com/acbodine/iot/hub/platform"
)

func main() {

    // TODO: Make this changed based on paramters.
    // Get an Bluemix Platformer.
    platformer, err := platform.Factory(platform.Bluemix)
    if err != nil {
        panic(err)
    }

    pumped := make(chan *event.Event)

    // Tell Platformer to start pumping Events.
    if err = platformer.Pump(pumped); err != nil {
        panic(err)
    }
    defer func () {
        if err := platformer.Terminate(); err != nil {
            panic(err)
        }
    }()

    // Create a MemoryEventer, and register with Platformer.
    eventer := local.MemoryEventer(time.Minute)
    if err = platformer.Register(eventer); err != nil {
        panic(err)
    }

    var wg sync.WaitGroup

    wg.Add(1)
    go func () {
        defer wg.Done()

        // TODO: Make this interruptable.
        for {
            select {
            case <- pumped:
                // TODO: Make logging level configurable.
                // fmt.Println("Pumped event:", evt)
            }
        }
    }()

    wg.Wait()
}
