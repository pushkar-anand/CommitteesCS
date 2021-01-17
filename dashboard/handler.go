package dashboard

import (
	"committees/template"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	logger *logrus.Logger
}

func NewHandler(logger *logrus.Logger) *Handler {
	return &Handler{logger: logger}
}

func (h *Handler) Dashboard(w http.ResponseWriter, r *http.Request) {
	template.Render(w, "dashboard.html", nil)
}
