package render

import (
	"html/template"
)

func ParseTemplates() (*template.Template, error) {
	templateBuilder := template.New("")
	if t, _ := templateBuilder.ParseGlob("views/*/*.html"); t != nil {
		templateBuilder = t
	}

	return templateBuilder.ParseGlob("views/*.html")
}


