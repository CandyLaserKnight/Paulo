package session

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"reflect"
	"testing"
)

func TestSession_InitSession(t *testing.T) {

	p := &Session{
		CookieLifetime: "100",
		CookiePersist:  "true",
		CookieName:     "paulo",
		CookieDomain:   "localhost",
		SessionType:    "cookie",
	}

	var sm *scs.SessionManager

	ses := p.InitSession()

	var sesskind reflect.Kind
	var sesstype reflect.Type

	rv := reflect.ValueOf(ses)

	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		fmt.Println("for loop:", rv.Kind(), rv.Type(), rv)
		sesskind = rv.Kind()
		sesstype = rv.Type()

		rv = rv.Elem()
	}

	if !rv.IsValid() {
		t.Error("invalid type or kind; kind:", rv.Kind(), "type:", rv.Type())
	}

	if sesskind != reflect.ValueOf(sm).Kind() {
		t.Error("wrong kind returned resting cookie session. Expected", reflect.ValueOf(sm).Kind(), "and got", sesskind)
	}

	if sesstype != reflect.ValueOf(sm).Type() {
		t.Error("wrong type returned resting cookie session. Expected", reflect.ValueOf(sm).Type(), "and got", sesstype)
	}
}
