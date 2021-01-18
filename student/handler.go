package student

import (
	"committees/helpers"
	"committees/request"
	"committees/template"
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
		h.logger.WithError(err).Error("error adding faculty to DB")
		helpers.InternalError(w)
		return
	}

	http.Redirect(w, r, "/dashboard/students", http.StatusSeeOther)
}
