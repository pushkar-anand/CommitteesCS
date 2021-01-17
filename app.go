package main

import (
	"committees/db"
	"committees/faculty"
	"github.com/sirupsen/logrus"
	"net/http"

	"committees/dashboard"
	"github.com/gorilla/mux"
)

func addRoutes(r *mux.Router, logger *logrus.Logger) {
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	dbConn := db.GetDB(logger)

	dh := dashboard.NewHandler(logger)

	dr := r.PathPrefix("/dashboard").Subrouter()

	dashboard.AddRoutes(dr, dh)

	facultyRepo := faculty.NewRepository(dbConn)
	fh := faculty.NewHandler(logger, facultyRepo)
	fr := dr.PathPrefix("/faculty").Subrouter()

	faculty.AddRoutes(fr, fh)
}
