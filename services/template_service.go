package services

import (
	"path/filepath"
	"github.com/flosch/pongo2/v6"
)

// RenderTemplate loads, compiles, and executes a Pongo2 template.
// templateName should be relative to the "templates/" directory, e.g., "tpl.html".
func RenderTemplate(templateName string, data pongo2.Context) (string, error) {
	// In a real application, you might want to configure the base directory for templates
	// globally or pass it to this function, instead of hardcoding "templates".
	// For pongo2, if you set up a global template set with a base directory,
	// you might not need filepath.Join here and could just pass templateName.
	// pongo2.DefaultSet.SetBaseDirectory("templates") // Example global setting

	templatePath := filepath.Join("templates", templateName)
	tpl, err := pongo2.FromFile(templatePath)
	if err != nil {
		return "", err
	}
	return tpl.Execute(data)
}
