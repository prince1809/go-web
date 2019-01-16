package main

import (
	"html/template"
	"log"
	"os"
)

type Note struct {
	Title       string
	Description string
}

const tmpl = `Note - Title: {{ .Title }}, Description: {{ .Description }}`

func main() {
	// create an instance of Note struct
	note := Note{
		"text/templates",
		"Template generates textual output",
	}

	//create a new template with name
	t := template.New("note")

	// parse some content and generate a template
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	// Applies a parsed template to the data of Note object
	if err := t.Execute(os.Stdout, note); err != nil {
		log.Fatal("Execute: ", err)
		return
	}
}
