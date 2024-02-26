package render

import (
	"html/template"
)

func ParseTemplates() (*template.Template, error) {
	templateBuilder := template.New("")
	if t, _ := templateBuilder.ParseGlob("public/*/*.html"); t != nil {
		templateBuilder = t
	}

	return templateBuilder.ParseGlob("public/*.html")
}


