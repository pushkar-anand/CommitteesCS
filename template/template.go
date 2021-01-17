package template

import (
	"committees/config"
	"committees/helpers"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var templates *template.Template

func ParseTemplates() {
	templ := template.New("")
	err := filepath.Walk("./views", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = templ.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}

		return err
	})

	if err != nil {
		panic(err)
	}

	templates = templ
	return
}

func Render(w http.ResponseWriter, templateName string, data interface{}) {
	logger := config.GetLogger()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := templates.ExecuteTemplate(w, templateName, data)
	if err != nil {
		logger.WithError(err).Errorf("error rendering template: %s", templateName)
		helpers.InternalError(w)
	}
}
