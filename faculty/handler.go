package faculty

import (
	"committees/helpers"
	"committees/request"
	"committees/template"
	"encoding/csv"

	//"encoding/csv"
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

func (h *Handler) Faculty(w http.ResponseWriter, r *http.Request) {
	faculties, err := h.repository.FetchAll()
	if err != nil {
		h.logger.WithError(err).Error("error fetching faculties from DB")
		helpers.InternalError(w)
		return
	}

	data := map[string]interface{}{
		"Faculties": faculties,
	}

	template.Render(w, "faculty.html", data)
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	f := &Faculty{}

	ok := request.ReadFormDataAndValidate(w, r, f)
	if !ok {
		return
	}

	err := h.repository.Create(f)
	if err != nil {
		h.logger.WithError(err).Error("error adding faculty to DB")
		helpers.InternalError(w)
		return
	}

	http.Redirect(w, r, "/dashboard/faculty", http.StatusSeeOther)
}

func derefString(s *string) string {
	if s != nil {
		return *s
	}

	return ""
}

func (h *Handler) Csv(w http.ResponseWriter, r *http.Request) {
	faculties, err := h.repository.FetchAll()
	if err != nil {
		h.logger.WithError(err).Error("error fetching students from DB")
		helpers.InternalError(w)
		return
	}

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=Faculty.csv")
	wr := csv.NewWriter(w)
	_ = wr.Write([]string{"Name", "Designation"})

	for _, s := range faculties {
		d := make([]string, 0)
		d = append(d, derefString(s.Name))
		d = append(d, derefString(s.Designation))

		err = wr.Write(d)
		if err != nil {
			h.logger.WithError(err).Error("error writing csv response")
			helpers.InternalError(w)
			return
		}
	}

	wr.Flush()
}
