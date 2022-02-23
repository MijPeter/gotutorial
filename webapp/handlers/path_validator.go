package handlers

import (
	"errors"
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	matches := validPath.Match([]byte(r.URL.Path))
	if !matches {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}

	return validPath.FindStringSubmatch(r.URL.Path)[2], nil
}
