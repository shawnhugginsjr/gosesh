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

// SessionInterface implementation.
type session struct {
	id           string                 // ID of the session.
	attribute    map[string]interface{} // Constant attributes specified at session creation.
	setAttribute map[string]interface{} // Attributes stored in the session.
	created      time.Time              // Creation time.
	accessed     time.Time              // Last accessed time.
	timeout      time.Duration          // Session timeout.
}

// SessionInterface method implementation.
func (s *session) ID() string {
	return s.id
}

// SessionInterface method implementation.
func (s *session) Attribute(name string) (interface{}, bool) {
	elem, elemExists := s.attribute[name]
	return elem, elemExists
}

// SessionInterface method implementation.
func (s *session) Created() time.Time {
	return s.created
}

// SessionInterface method implementation.
func (s *session) Accessed() time.Time {
	return s.accessed
}

// SessionInterface method implementation.
func (s *session) Timeout() time.Duration {
	return s.timeout
}

// genID generates a random session id using the crypto/rand package.
func genID(length int) string {
	idBuf := make([]byte, length)
	io.ReadFull(rand.Reader, idBuf)
	return base64.URLEncoding.EncodeToString(idBuf)
}
