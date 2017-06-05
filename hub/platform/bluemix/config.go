package bluemix

import (
    "os"
)

// NOTE: Bluemix documentation for MQTT clients and gateways here:
// https://console.ng.bluemix.net/docs/services/IoT/gateways/mqtt.html#mqtt

type Config struct {
    OrgId       string
    TypeId      string
    DeviceId    string
    AuthToken   string
}

func ConfigFromEnvironment() (*Config) {
    c := &Config{}

    c.OrgId        = os.Getenv("WATSON_IOT_ORGID")
    c.TypeId       = os.Getenv("WATSON_IOT_TYPEID")
    c.DeviceId     = os.Getenv("WATSON_IOT_DEVICEID")
    c.AuthToken    = os.Getenv("WATSON_IO_AUTHTOKEN")

    return c
}
