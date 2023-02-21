package session

import (
	"github.com/alexedwards/scs/v2"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Session struct {
	CookieLifetime string
	CookiePersist  string
	CookieName     string
	CookieDomain   string
	SessionType    string
	CookieSecure   string
}

func (p *Session) InitSession() *scs.SessionManager {
	var persist, secure bool

	// how log should sessions last?
	minutes, err := strconv.Atoi(p.CookieLifetime)
	if err != nil {
		minutes = 60
	}

	// should cookies persist?
	if strings.ToLower(p.CookiePersist) == "true" {
		persist = true
	}

	// must cookies be secure?
	if strings.ToLower(p.CookieSecure) == "true" {
		secure = true
	}

	// create session
	session := scs.New()
	session.Lifetime = time.Duration(minutes) * time.Minute
	session.Cookie.Persist = persist
	session.Cookie.Name = p.CookieName
	session.Cookie.Secure = secure
	session.Cookie.Domain = p.CookieDomain
	session.Cookie.SameSite = http.SameSiteLaxMode

	// which session store?
	switch strings.ToLower(p.SessionType) {
	case "redis":

	case "mysql", "mariadb":

	case "postgres", "postgresql":

	default:
		// cookie
	}

	return session
}
