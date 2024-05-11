package handlers

import (
	"log"
	"net/http"

	"github.com/Akihira77/fullstack-go-htmx/components"
	"github.com/Akihira77/fullstack-go-htmx/views"
)

func (h *Handler) HandleCarsList(w http.ResponseWriter, r *http.Request) {
	cars, err := h.store.GetCars()
	if err != nil {
		log.Println(err)
		return
	}

	views.Cars(cars).Render(r.Context(), w)
}

func (h *Handler) HandleDeleteCar(w http.ResponseWriter, r *http.Request) {
	err := h.store.DeleteCar(r.PathValue("id"))
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) HandleCarSearch(w http.ResponseWriter, r *http.Request) {
	cars, err := h.store.FindCarsByNameMakeOrBrand(r.FormValue("search"))
	if err != nil {
		log.Println(err)
		return
	}

	components.CarsList(cars).Render(r.Context(), w)
}
