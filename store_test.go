package gosesh

import (
	"testing"
	"time"
)

func TestNewMemStore(t *testing.T) {
	var store StoreInterface
	// Check if MemStore impliments Store Interface
	memStore := NewMemStore(time.Minute)
	defer memStore.Close()
	store = memStore
	if store != memStore {
		t.Error("The stores are not the same")
	}
}

func TestAddSession(t *testing.T) {
	session := NewSession(10, time.Minute)
	memStore := NewMemStore(time.Minute)
	defer memStore.Close()

	memStore.Add(session)
	if memStore.sessions[session.ID()] != session {
		t.Error("A Session failed to be added to the memStore")
	}
}

func TestGetSession(t *testing.T) {
	session := NewSession(10, time.Minute)
	memStore := NewMemStore(time.Minute)
	defer memStore.Close()

	_, err := memStore.Get("invalid-id")
	if err == nil {
		t.Error("An error was not returned when using an invalid Session ID")
	}

	memStore.sessions[session.ID()] = session
	sess, err := memStore.Get(session.ID())
	if err != nil {
		t.Error("A session was not retrieved")
	}
	if sess.ID() != session.ID() {
		t.Error("A different Session was retrieved than saved.")
	}
}

func TestRemoveSession(t *testing.T) {
	session := NewSession(10, time.Minute)
	memStore := NewMemStore(time.Minute)
	defer memStore.Close()

	memStore.sessions[session.ID()] = session
	memStore.Remove(session)
	_, exists := memStore.sessions[session.ID()]
	if exists != false {
		t.Error("The Session was not removed.")
	}
}

func TestClearTimeouts(t *testing.T) {
	var sessionInterval time.Duration
	sessionCount := 6
	memStore := NewMemStore(100 * time.Millisecond)

	for i := 1; i <= sessionCount; i++ {
		if i%2 == 0 {
			sessionInterval = 20 * time.Millisecond
		} else {
			sessionInterval = 900 * time.Millisecond
		}
		sess := NewSession(2, sessionInterval)
		memStore.Add(sess)
	}

	time.Sleep(200 * time.Millisecond)
	if len(memStore.sessions) != (sessionCount / 2) {
		t.Error("Sessions were not deleted during timeout sweep.")
	}
}
