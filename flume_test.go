package utils

import (
	"testing"
)

var (
	flumeHost     = "192.168.1.110"
	flumePort int = 60000
)

func TestAppend(t *testing.T) {
	cl := NewFlumeClient(flumeHost, flumePort)
	evt := NewThriftFlumeEvent()
	evt.Headers = map[string]string{
		"type": "online",
	}
	evt.Body = []byte("Test body")
	err := cl.Append(evt)
	if err != nil {
		t.Error(err)
	}
}
