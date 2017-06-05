package bluemix

import (
    "fmt"

    MQTT "github.com/eclipse/paho.mqtt.golang"
)

// NOTE: Bluemix documentation for MQTT clients and gateways here:
// https://console.ng.bluemix.net/docs/services/IoT/gateways/mqtt.html#mqtt

type client struct {
    config      *Config
    mqttClient  MQTT.Client
}

// Connect authenticates to the Watson Iot Platform with the
// provided configuration, and opens a new MQTT client connection
// to the corresponding MQTT hub.
func Connect(c *Config) (*client, error) {
    if c == nil {
        c = ConfigFromEnvironment()
    }

    ident := fmt.Sprintf(
        "g:%s:%s:%s",
        c.OrgId,
        c.TypeId,
        c.DeviceId,
    )

    opts := MQTT.NewClientOptions()

    broker := fmt.Sprintf(
        "tcp://%s.messaging.internetofthings.ibmcloud.com:1883",
        c.OrgId,
    )
    opts.AddBroker(broker)

    opts.SetClientID(ident)
    opts.SetUsername("use-token-auth")
    opts.SetPassword(c.AuthToken)

    opts.SetAutoReconnect(false)

    cli := MQTT.NewClient(opts)
    if t := cli.Connect(); t.Wait() && t.Error() != nil {
        return nil, t.Error()
    }

    client := &client{
        config:     c,
        mqttClient: cli,
    }

    return client, nil
}

// Publish takes a []byte payload and publishes it
// to the underlying MQTT connection.
func (c *client) Publish(data []byte) error {
    topic := fmt.Sprintf(
        "iot-2/type/%s/id/%s/evt/usage/fmt/json",
        c.config.TypeId,
        c.config.DeviceId,
    )

    t := c.mqttClient.Publish(topic, 0, false, data)
    t.Wait()

    if t.Error() != nil {
        return t.Error()
    }

    return nil
}

func (c *client) Disconnect() error {
    c.mqttClient.Disconnect(250)

    return nil
}
