package message

import "testing"

func TestGetMessage(t *testing.T) {
	message := GetMessage()
	if message != "Halo, Golang!" {
		t.Errorf("Message is not valid: %s", message)
	}
}