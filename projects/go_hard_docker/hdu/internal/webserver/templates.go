package webserver

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// --- составление списка шаблонов(обход всего каталога)
func makeTemplatesList(root_path string) []string {

	var all_templates []string

	filepath.Walk(root_path, func(path string, info os.FileInfo, err error) error {
		filename := info.Name()
		if strings.HasSuffix(filename, ".html") {
			fmt.Println(path)
			all_templates = append(all_templates, path)
		}
		return nil
	})
	return all_templates
}
