package committee

import (
	"github.com/gorilla/mux"
	"net/http"
)

func AddRoutes(r *mux.Router, h *Handler) {
	r.HandleFunc("/", h.Committee).Methods(http.MethodGet)
	r.HandleFunc("/add", h.Add).Methods(http.MethodPost)
	r.HandleFunc("/{committee_id}/faculty/{faculty_id}", h.AddFaculty).Methods(http.MethodPost)
}
