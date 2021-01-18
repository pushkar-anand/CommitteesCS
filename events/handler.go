package events

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

func (h *Handler) Event(w http.ResponseWriter, r *http.Request) {
	events, err := h.repository.FetchAll()
	if err != nil {
		h.logger.WithError(err).Error("error fetching events from DB")
		helpers.InternalError(w)
		return
	}

	data := map[string]interface{}{
		"Events": events,
	}

	template.Render(w, "events.html", data)
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	s := &Event{}

	ok := request.ReadFormDataAndValidate(w, r, s)
	if !ok {
		return
	}

	err := h.repository.Create(s)
	if err != nil {
		h.logger.WithError(err).Error("error adding events to DB")
		helpers.InternalError(w)
		return
	}

	http.Redirect(w, r, "/dashboard/events", http.StatusSeeOther)
}
