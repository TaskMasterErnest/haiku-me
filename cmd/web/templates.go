package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"github.com/TaskMasterErnest/internal/models"
	"github.com/TaskMasterErnest/ui"
)

type templateData struct {
	CurrentYear     int
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

func humanDate(t time.Time) string {
	//return empty string if time has the zero value
	if t.IsZero() {
		return ""
	}

	//convert time into UTC before formatting
	return t.UTC().Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	//use fs.Glob() to get a slice of all filepaths in the ui.Files embedded filesystem
	pages, err := fs.Glob(ui.Files, "html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		//the filepath patterns for the templates to parse
		patterns := []string{
			"html/base.tmpl.html",
			"html/partials/*.html",
			page,
		}

		//use ParseFS instead of ParseFiles to parse the template files from the ui.Files embedded system
		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil

}
