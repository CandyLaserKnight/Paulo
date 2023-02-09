package handlers

import (
	"github.com/candylaserknight/paulo"
	"net/http"
)

type Handlers struct {
	App *paulo.Paulo
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}
