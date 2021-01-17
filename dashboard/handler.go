package dashboard

import(
	"net/http"
	"committees/template"
)

type Handler struct {
	
}

func (h *Handler) Dashboard(w http.ResponseWriter, r *http.Request) {
	err := template.Render(w, "dashboard.html", nil)
	if err != nil {
		panic(err)
	}
}
