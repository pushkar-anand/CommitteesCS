package main

import (
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/gorilla/mux"
	"committees/dashboard"
)

func addRoutes(r *mux.Router, logger *logrus.Logger) {
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	dh := &dashboard.Handler{}
	dashboard.AddRoutes(r, dh)
}
