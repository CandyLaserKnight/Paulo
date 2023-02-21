package paulo

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (p *Paulo) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	if p.Debug {
		mux.Use(middleware.Logger)
	}
	mux.Use(middleware.Recoverer)
	mux.Use(p.SessionLoad)

	return mux
}
