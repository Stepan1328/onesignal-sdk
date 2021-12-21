package onesignal

import (
	"encoding/json"
	"log"
	"strconv"
)

const (
	methodKey  = "method"
	keyTypeKey = "key_type"
	payloadKey = "payload"
)

// Params represents a set of parameters that gets passed to a request.
type Params map[string]Param

type Param struct {
	URLValue string
	Payload  []byte

	SpecificValue string
}

func NewParamsWithMethod(httpMethod string) Params {
	params := make(Params, 1)
	params[methodKey] = Param{
		SpecificValue: httpMethod,
	}

	return params
}

// Method return string HTTP method to send data
func (p Params) Method() string {
	return p[methodKey].SpecificValue
}

const (
	typeRestAPI = "rest_api_type"
	typeAuth    = "auth_type"
)

func (p Params) SetKeyType(keyType string) Params {
	p[keyTypeKey] = Param{
		SpecificValue: keyType,
	}

	return p
}

func (p Params) KeyType() string {
	return p[keyTypeKey].SpecificValue
}

// AddJSONPayload adds a non-empty value that is encoded in json.
func (p Params) AddJSONPayload(value interface{}) {
	if value != "" {
		data, err := json.Marshal(value)
		if err != nil {
			log.Fatalf("unexpected marshaling error: %s\n", err.Error())
		}

		p[payloadKey] = Param{
			Payload: data,
		}
	}
}

// GetPayload returns data to send the payload over http
func (p Params) GetPayload() []byte {
	return p[payloadKey].Payload
}

// AddURLNonEmpty adds an url value if it not an empty string.
func (p Params) AddURLNonEmpty(key, value string) {
	if value != "" {
		p[key] = Param{
			URLValue: value,
		}
	}
}

// AddURLNonZero adds an url value if it is not zero.
func (p Params) AddURLNonZero(key string, value int) {
	if value != 0 {
		p[key] = Param{
			URLValue: strconv.Itoa(value),
		}
	}
}
