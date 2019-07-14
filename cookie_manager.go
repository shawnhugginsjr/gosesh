package gosesh

import (
	"fmt"
	"net/http"
	"time"
)

// CookieManager manages session access through cookies.
type CookieManager struct {
	Store             StoreInterface
	SessionCookieName string        // Name of cookies that stores the session id
	CookiePath        string        // Path of the cookie
	CookieMaxAge      time.Duration // Max age for the cookie
	Secure            bool          // If session cookies can only be sent over HTTPS
}

// NewCookieManager returns a pointer to a CookieManager
func NewCookieManager(store StoreInterface, sessionCookieName, cookiePath string, cookieMaxAge time.Duration, secure bool) *CookieManager {
	return &CookieManager{
		Store:             store,
		SessionCookieName: sessionCookieName,
		CookiePath:        cookiePath,
		CookieMaxAge:      cookieMaxAge,
		Secure:            secure,
	}
}

// Get attemps to retrieves a session using the session cookie of a request.
func (c *CookieManager) Get(r *http.Request) (SessionInterface, error) {
	cookie, err := r.Cookie(c.SessionCookieName)
	if err != nil {
		return nil, fmt.Errorf("Cookie %s was not in request", c.SessionCookieName)
	}

	return c.Store.Get(cookie.Value)
}

// Add adds a session to the store and attaches a session cookie to the response.
func (c *CookieManager) Add(session SessionInterface, w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:     c.SessionCookieName,
		Value:    session.ID(),
		Path:     c.CookiePath,
		MaxAge:   int(c.CookieMaxAge.Seconds()),
		Secure:   c.Secure,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
	c.Store.Add(session)
}

// Remove removes the session cookie from the client and removes the session from the store.
func (c *CookieManager) Remove(session SessionInterface, w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:     c.SessionCookieName,
		Value:    "",
		Path:     c.CookiePath,
		MaxAge:   -1,
		Secure:   c.Secure,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
	c.Store.Remove(session)
}
