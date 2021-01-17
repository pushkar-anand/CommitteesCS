package dashboard

import "github.com/gorilla/mux"

func AddRoutes(r *mux.Router, h *Handler) {
	r.HandleFunc("/dashboard", h.Dashboard)
}
