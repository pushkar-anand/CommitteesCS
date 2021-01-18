package committee

import (
	"committees/helpers"
	"committees/request"
	"committees/template"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Handler struct {
	logger     *logrus.Logger
	repository *Repository
}

func NewHandler(logger *logrus.Logger, repository *Repository) *Handler {
	return &Handler{logger: logger, repository: repository}
}

func (h *Handler) Committee(w http.ResponseWriter, r *http.Request) {
	committees, err := h.repository.FetchAll()
	if err != nil {
		h.logger.WithError(err).Error("error fetching committees from DB")
		helpers.InternalError(w)
		return
	}

	data := map[string]interface{}{
		"Committees": committees,
	}

	template.Render(w, "committee.html", data)
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	c := &Committee{}

	ok := request.ReadFormDataAndValidate(w, r, c)
	if !ok {
		return
	}

	err := h.repository.Create(c)
	if err != nil {
		h.logger.WithError(err).Error("error adding committee to DB")
		helpers.InternalError(w)
		return
	}

	http.Redirect(w, r, "/dashboard/committees", http.StatusSeeOther)
}

func (h *Handler) AddFaculty(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	committeeID, err := strconv.ParseUint(params["committee_id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = fmt.Fprintf(w, "Invalid committe id: %s", params["committee_id"])
		return
	}

	facultyID, err := strconv.ParseUint(params["faculty_id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = fmt.Fprintf(w, "Invalid faculty id: %s", params["faculty_id"])
		return
	}

	err = h.repository.AddFacultyToCommittee(uint(committeeID), uint(facultyID))
	if err != nil {
		h.logger.WithError(err).Error("error adding faculty to committee")
		helpers.InternalError(w)
		return
	}

	http.Redirect(w, r, "/dashboard/committees", http.StatusSeeOther)
}
