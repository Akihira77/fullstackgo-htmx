package handlers

import "github.com/Akihira77/fullstack-go-htmx/store"

type Handler struct {
	store *store.Storage
}

func New(store *store.Storage) *Handler {
	return &Handler{
		store: store,
	}
}
