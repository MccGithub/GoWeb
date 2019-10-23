package view_page

import (
	"html/template"
	"io"
)

//type Data map[string]interface{}

//func View(file string, w io.Writer, data Data) error {
func View(file string, w io.Writer, data interface{}) error {
	temp, err := template.ParseFiles(getTempPath(file))
	if err != nil {
		return err
	}

	return temp.Execute(w, data)
}
