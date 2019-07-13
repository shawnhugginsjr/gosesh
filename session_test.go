package gosesh

import (
	"testing"
	"time"
)

func TestNewSession(t *testing.T) {
	session := NewSession(16, 2*time.Minute)
	// Check if session impliments SessionInterface
	var sesh SessionInterface = session

	if session.created != sesh.Accessed() {
		t.Error("Session's initial created and accessed times are different.")
	}
}

func TestSessionAttributes(t *testing.T) {
	session := NewSession(16, 2*time.Minute)

	if len(session.attributes) != 0 {
		t.Error("New Session Attributes length are not initialized to zero.")
	}

	name := "john"
	session.SetAttribute("name", name)
	elem, exists := session.Attribute("name")
	if !exists {
		t.Error("An attribute was not succusfly saved into the Sessions Attributes.")
	}

	if elem != name {
		t.Error("A key returned a value different than what was saved in the Session.")
	}
}
