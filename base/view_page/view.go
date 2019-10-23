package view_page

import (
	"html/template"
	"io"
)

type data map[string]interface{}

func View(file string, w io.Writer, d data) error {
	temp, err := template.ParseFiles(getPath(file))
	if err != nil {
		return err
	}

	return temp.Execute(w, d)
}
