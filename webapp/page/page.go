package page

import (
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

var PATH = "PagesDb"

func (p *Page) Save() error {
	filename := PATH + "/" + p.Title
	return os.WriteFile(filename, p.Body, 0600)
}

func Load(title string) (*Page, error) {
	filename := PATH + "/" + title
	body, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}
