package handlers

import (
	"net/http"

	"github.com/Akihira77/fullstack-go-htmx/views"
)

func (h *Handler) HandleHome(w http.ResponseWriter, r *http.Request) {
	views.Home().Render(r.Context(), w)
}
