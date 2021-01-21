package committee

import (
	"committees/helpers"
	"committees/request"
	"committees/template"
	"encoding/csv"
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

	faculties, err := h.repository.FetchFaculties()
	if err != nil {
		h.logger.WithError(err).Error("error fetching faculties from DB")
		helpers.InternalError(w)
		return
	}

	data := map[string]interface{}{
		"Committees": committees,
		"Faculties":  faculties,
	}

	template.Render(w, "committee.html", data)
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	c := &Committee{}

	ok := request.ReadJSONAndValidate(w, r, c)
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

func derefString(s *string) string {
	if s != nil {
		return *s
	}

	return ""
}

func (h *Handler) Csv(w http.ResponseWriter, r *http.Request) {
	committees, err := h.repository.FetchAll()
	if err != nil {
		h.logger.WithError(err).Error("error fetching students from DB")
		helpers.InternalError(w)
		return
	}

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=Committees.csv")
	wr := csv.NewWriter(w)
	_ = wr.Write([]string{"Name", "Description", "Members", "Created At"})

	for _, s := range committees {
		d := make([]string, 0)
		d = append(d, derefString(s.Name))
		d = append(d, derefString(s.Description))

		member := ""

		for _, m := range s.Members {
			//member = append(member, ", "+*m.Name)
			member = member + ",\n " + *m.Name
		}

		d = append(d, member)
		d = append(d, s.CreationDate.Format("Jan 02, 2006"))

		err = wr.Write(d)
		if err != nil {
			h.logger.WithError(err).Error("error writing csv response")
			helpers.InternalError(w)
			return
		}
	}

	wr.Flush()
}

