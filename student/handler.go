package student

import (
	"committees/helpers"
	"committees/request"
	"committees/template"
	"encoding/csv"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	logger     *logrus.Logger
	repository *Repository
}

func NewHandler(logger *logrus.Logger, repository *Repository) *Handler {
	return &Handler{logger: logger, repository: repository}
}

func (h *Handler) Student(w http.ResponseWriter, r *http.Request) {
	students, err := h.repository.FetchAll()
	if err != nil {
		h.logger.WithError(err).Error("error fetching students from DB")
		helpers.InternalError(w)
		return
	}

	data := map[string]interface{}{
		"Students": students,
	}

	template.Render(w, "students.html", data)
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	s := &Student{}

	ok := request.ReadFormDataAndValidate(w, r, s)
	if !ok {
		return
	}

	err := h.repository.Create(s)
	if err != nil {
		h.logger.WithError(err).Error("error adding students to DB")
		helpers.InternalError(w)
		return
	}

	http.Redirect(w, r, "/dashboard/students", http.StatusSeeOther)
}

func derefString(s *string) string {
	if s != nil {
		return *s
	}

	return ""
}

func (h *Handler) Csv(w http.ResponseWriter, r *http.Request) {
	students, err := h.repository.FetchAll()
	if err != nil {
		h.logger.WithError(err).Error("error fetching students from DB")
		helpers.InternalError(w)
		return
	}

	data := make([][]string, 0)

	data = append(data, []string{"Name", "Email", "Phone", "USN"})

	for _, s := range students {
		d := make([]string, 0)
		d = append(d, derefString(s.Name))
		d = append(d, derefString(s.Email))
		d = append(d, derefString(s.Phone))
		d = append(d, derefString(s.Usn))

		data = append(data, d)
	}

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=TheCSVFileName.csv")
	wr := csv.NewWriter(w)
	err = wr.WriteAll(data)
	if err != nil {
		h.logger.WithError(err).Error("error writing csv response")
		helpers.InternalError(w)
		return
	}

	wr.Flush()
}
