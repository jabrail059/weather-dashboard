package core_http_session

import (
	"crypto/rand"
	"net/http"
	"time"
)

func Cookie() *http.Cookie {
	value := rand.Text()
	return &http.Cookie{
		Name:     "session_id",
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24 * 365),
		SameSite: http.SameSiteStrictMode,
	}
}

func GetOrCreate(r *http.Request) (*http.Cookie, bool) {
	cookie, err := r.Cookie("session_id")
	if err == nil {
		return cookie, false
	}
	return Cookie(), true
}
