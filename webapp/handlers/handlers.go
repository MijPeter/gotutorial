package handlers

import (
	"errors"
	"html/template"
	"net/http"
	"os"

	"github.com/MijPeter/gotutorial/tree/main/webapp/page"
)

var templates = template.Must(template.ParseFiles("templates/view.html", "templates/edit.html"))

func ViewHandler(w http.ResponseWriter, r *http.Request, filename string) {
	p, err := page.Load(filename)

	if errors.Is(err, os.ErrNotExist) {
		http.Redirect(w, r, "/edit/"+filename, http.StatusFound)
		return
	} else if err != nil {
		serverError(w, err)
		return
	}

	renderTemplate(w, "view.html", p)
}

func EditHandler(w http.ResponseWriter, r *http.Request, filename string) {
	p, err := page.Load(filename)

	if errors.Is(err, os.ErrNotExist) {
		p = &page.Page{Title: filename}
	} else if err != nil {
		serverError(w, err)
		return
	}

	renderTemplate(w, "edit.html", p)
}

func SaveHandler(w http.ResponseWriter, r *http.Request, filename string) {
	body := []byte(r.FormValue("body"))

	p := page.Page{Title: filename, Body: body}
	err := p.Save()

	if err != nil {
		serverError(w, err)
		return
	}

	http.Redirect(w, r, "/view/"+filename, http.StatusFound)
}

func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filename, err := getTitle(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fn(w, r, filename)
	}
}

func serverError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func renderTemplate(w http.ResponseWriter, templateName string, p *page.Page) {
	err := templates.ExecuteTemplate(w, templateName, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
