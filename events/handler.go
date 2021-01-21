package events

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

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=Events.csv")
	wr := csv.NewWriter(w)
	_ = wr.Write([]string{"Name", "Start Date", "End Date", "USN"})

	for _, s := range students {
		d := make([]string, 0)
		d = append(d, derefString(s.Name))
		d = append(d, s.StartDate.Format("Jan 02, 2006"))
		d = append(d, s.EndDate.Format("Jan 02, 2006"))
		d = append(d, derefString(s.TotalExpenditure))

		err = wr.Write(d)
		if err != nil {
			h.logger.WithError(err).Error("error writing csv response")
			helpers.InternalError(w)
			return
		}
	}

	wr.Flush()
}

