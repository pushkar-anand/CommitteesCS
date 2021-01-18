package events

import (
	"github.com/gorilla/mux"
	"net/http"
)

func AddRoutes(r *mux.Router, h *Handler) {
	r.HandleFunc("/", h.Event).Methods(http.MethodGet)
	r.HandleFunc("/add", h.Add).Methods(http.MethodPost)
}
