package main

import (
	"fmt"
	"os"
)

var folder string = "pages"

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title
	return os.WriteFile(filename, p.Body, 0600)
}

func load(title string) (*Page, error) {
	filename := title
	body, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func main() {
	page := Page{Title: "webapptest", Body: []byte{120, 121, 123, 125, '3', '1'}}
	fmt.Println("Saving file...")
	page.save()
	fmt.Println("Loading file...")

	loadedPage, _ := load("webapptest")
	fmt.Println(string(loadedPage.Body))
}
