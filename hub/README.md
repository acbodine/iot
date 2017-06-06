# hub

## Overview

### Platformers

Platformers take care of shipping data of interest from Eventers to various
hosted IoT platforms. This class serves to abstract the specifics of data
transmission for each platform. Many Eventers can send to the same Platformer.

#### Local
TODO

#### Bluemix
TODO

### Eventers

Eventers have the role of aggregating data from various IoT devices and making
it available in an orderly manner to the Platformers. This should be treated
as an adapter interface, and you will need to implement an Eventer for each
type of event you want to capture at the hub.

#### Local
- Mock
- Memory

## Build

Raspberry Pi3
```
$ export GOOS=linux
$ export GOARCH=arm
$ go build -o hub/hub github.com/acbodine/iot/hub
```
