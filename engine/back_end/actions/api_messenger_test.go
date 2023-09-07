package actions

import "testing"

func TestGetMessengerActions(t *testing.T) {
	messengerActions := GetMessengerActions()
	if len(messengerActions) == 0 {
		t.Error("No messenger actions found")
	}
}

func TestNewMessenger(t *testing.T) {
	m := NewMessenger()
	if m == nil {
		t.Error("NewMessenger() should not return nil")
	}
}

func TestSendAction(t *testing.T) {
	messengerActions := GetMessengerActions()
	send := messengerActions[0]
	if send.Identifier != "quando.message.send" {
		t.Error("Identifier should be quando.message.send")
	}
	if send.Type != "action" {
		t.Error("Type should be action")
	}
}
