package gosesh

import (
	"fmt"
	"sync"
)

// StoreInterface is responsible for fetching and saving sessions by their ID.
type StoreInterface interface {
	// Get returns a pointer to a SessionInterface by its id.
	// The returned session will have an updated access time set to the current time.
	// An error is returned if this store does not contain a session with the specified id.
	Get(id string) (SessionInterface, error)

	// Add adds a new SessionInterface to the store.
	Add(session SessionInterface)

	// Remove removes a SessionInterface from the store.
	Remove(session SessionInterface)
}

// MemStore is a StoreInterface implementation.
type MemStore struct {
	sessions map[string]SessionInterface // Map of sessions (mapped from ID)
	mux      *sync.RWMutex               // mutex to synchronize access to sessions
}

// NewMemStore returns a pointer to a MemStore
func NewMemStore() *MemStore {
	s := &MemStore{
		sessions: make(map[string]SessionInterface),
		mux:      &sync.RWMutex{},
	}

	return s
}

// Get implements Store.Get().
func (s *MemStore) Get(id string) (SessionInterface, error) {
	s.mux.RLock()
	defer s.mux.RUnlock()

	session, sessionExists := s.sessions[id]
	if !sessionExists {
		return nil, fmt.Errorf("Session ID %s does not exist", id)
	}

	session.Access()
	return session, nil
}

// Add implements Store.Add().
func (s *MemStore) Add(session SessionInterface) {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.sessions[session.ID()] = session
}

// Remove implements Store.Remove().
func (s *MemStore) Remove(session SessionInterface) {
	s.mux.Lock()
	defer s.mux.Unlock()

	delete(s.sessions, session.ID())
}
