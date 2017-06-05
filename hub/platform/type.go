package platform

import (
    "bytes"
    "encoding/json"
)

type Type int

const (
    Local Type = iota
    Bluemix

    // TODO: Add more platform types here
)

var Type2String = map[Type]string{
    Local:      "local",
    Bluemix:    "bluemix",
}

var String2Type = map[string]Type{
    "local":    Local,
    "bluemix":  Bluemix,
}

// Implement fmt.Stringer interface.
func (t Type) String() string {
    return Type2String[t]
}

// Implement json.Marshaler interface.
func (t Type) MarshalJSON() ([]byte, error) {
    b := bytes.NewBufferString(`"`)
    b.WriteString(Type2String[t])
    b.WriteString(`"`)

    return b.Bytes(), nil
}

// Implement json.Unmarshaler interface.
func (t Type) UnmarshalJSON(b []byte) error {
    var s string

    if err := json.Unmarshal(b, &s); err != nil {
        return err
    }

    t = String2Type[s]

    return nil
}
