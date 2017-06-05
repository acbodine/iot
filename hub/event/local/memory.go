package local

import (
    "fmt"
    "encoding/json"
    "runtime"
    "sync"
    "time"

    "github.com/acbodine/iot/hub/event"
)

const (
    wait = time.Second * 10
)

type memory struct {
    events      chan *event.Event
    terminate   chan bool

    stats       *runtime.MemStats
    timeout     time.Duration

    sync.WaitGroup
}

func MemoryEventer(timeout time.Duration) event.Eventer {
    if timeout == 0 {
        timeout = wait
    }

    return &memory{
        terminate:  make(chan bool),
        timeout:    timeout,

        stats:      &runtime.MemStats{},
    }
}

// Implement the event.Eventer interface.
func (m *memory) Emit(events chan *event.Event) error {
    m.events = events

    m.Add(1)
    go m.emit()

    return nil
}

func (m *memory) emit() {
    defer m.Done()

    for {
        timeout := time.After(m.timeout)

        select {

        case <- m.terminate:
            return

        case <- timeout:
            runtime.ReadMemStats(m.stats)

            mem := map[string]interface{}{
                "Alloc": fmt.Sprintf("%v", m.stats.Alloc / 1024),
                "TotalAlloc": fmt.Sprintf("%v", m.stats.TotalAlloc / 1024),
                "Sys": fmt.Sprintf("%v", m.stats.Sys / 1024),
                "NumGC": fmt.Sprintf("%v", m.stats.NumGC),
            }

            payload := map[string]interface{}{
                "memory": mem,
            }

            data, err := json.Marshal(payload)
            if err != nil {
                fmt.Println(err)
            }

            m.events <- event.New(data, time.Now())

            break

        }
    }
}

// Implement the event.Eventer interface.
func (m *memory) Terminate() error {
    m.terminate <- true
    m.Wait()

    m.events = nil

    return nil
}
