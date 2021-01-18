package main

import (
	"committees/committee"
	"committees/db"
	"committees/events"
	"committees/faculty"
	"committees/student"
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

	//	Student endpoint (comment can be removed latter)
	studentRepo := student.NewRepository(dbConn)
	sh := student.NewHandler(logger, studentRepo)
	sr := dr.PathPrefix("/students").Subrouter()

	student.AddRoutes(sr, sh)

	committeeRepo := committee.NewRepository(dbConn)
	ch := committee.NewHandler(logger, committeeRepo)
	cr := dr.PathPrefix("/committees").Subrouter()

	committee.AddRoutes(cr, ch)

	eventsRepo := events.NewRepository(dbConn)
	eh := events.NewHandler(logger, eventsRepo)
	er := dr.PathPrefix("/events").Subrouter()

	events.AddRoutes(er, eh)
}
