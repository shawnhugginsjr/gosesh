package gosesh

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"time"
)

// SessionInterface stores and retrieves variable attributes.
type SessionInterface interface {
	// ID returns the id of the session.
	ID() string

	// Attribute returns the value of an attribute stored in the session and if it exists.
	Attribute(name string) (interface{}, bool)

	// SetAttribute Sets the value of an attribute stored in the session.
	SetAttribute(name string, value interface{})

	// Created returns the session creation time.
	Created() time.Time

	// Accessed returns the time when the session was last accessed.
	Accessed() time.Time

	// Timeout returns the session timeout.
	// A session may be removed automatically if it is not accessed for this duration.
	Timeout() time.Duration

	// Access records an access to the session by updating its last accessed time to
	// the current time.
	Access()
}

// Session is a SessionInterface implementation.
type Session struct {
	id        string                 // ID of the session.
	attribute map[string]interface{} // attributes specified at session creation.
	created   time.Time              // Creation time.
	accessed  time.Time              // Last accessed time.
	timeout   time.Duration          // Session timeout.
}

// NewSession returns a pointer to a new session.
func NewSession(idByteLength int, sessionTimeout time.Duration) *Session {
	currentTime := time.Now()
	sess := Session{
		id:        genID(idByteLength),
		attribute: make(map[string]interface{}),
		created:   currentTime,
		accessed:  currentTime,
		timeout:   sessionTimeout,
	}

	return &sess
}

// ID is SessionInterface method implementation.
func (s *Session) ID() string {
	return s.id
}

// Attribute is SessionInterface method implementation.
func (s *Session) Attribute(name string) (interface{}, bool) {
	elem, elemExists := s.attribute[name]
	return elem, elemExists
}

// Created is SessionInterface method implementation.
func (s *Session) Created() time.Time {
	return s.created
}

// Accessed is SessionInterface method implementation.
func (s *Session) Accessed() time.Time {
	return s.accessed
}

// Timeout is SessionInterface method implementation.
func (s *Session) Timeout() time.Duration {
	return s.timeout
}

// genID generates a random session id using the crypto/rand package.
func genID(length int) string {
	idBuf := make([]byte, length)
	io.ReadFull(rand.Reader, idBuf)
	return base64.URLEncoding.EncodeToString(idBuf)
}
