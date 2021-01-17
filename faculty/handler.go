package faculty

import (
	"committees/helpers"
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

func (h *Handler) Faculty(w http.ResponseWriter, r *http.Request) {
	faculties, err := h.repository.FetchAll()
	if err != nil {
		h.logger.WithError(err).Error("error fetching faculties from DB")
		helpers.InternalError(w)
		return
	}

	template.Render(w, "faculty.html", faculties)

}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {

}
