package templates

import (
	"bytes"
	"html/template"
	"io"

	"github.com/hysem/charts/config"
	"github.com/labstack/echo/v4"
)

//Template holds template rendering information
type Template struct {
	templates *template.Template
}

// Current holds application template
var Current *Template

//Render reders a template
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

//Execute executes template and data
func (t *Template) Execute(name string, data interface{}) ([]byte, error) {
	var tpl bytes.Buffer
	if err := t.templates.ExecuteTemplate(&tpl, name, data); err != nil {
		return nil, err
	}
	return tpl.Bytes(), nil
}

//Set sets the template for the application
func Set(e *echo.Echo) {
	Current = &Template{
		templates: template.Must(template.New("tpl").Delims("[[", "]]").ParseGlob(config.Current.Path.Template + "/*.html")),
	}
	e.Renderer = Current
}

//DefaultParams retruns a list of default params to be used in template
func DefaultParams(c echo.Context) map[string]interface{} {
	defaultParams := map[string]interface{}{
		"AppURL":    config.Current.URL.App,
		"AppName":   config.Current.AppName,
		"PageTitle": config.Current.AppName,
		"AssetURL":  config.Current.URL.Asset,
	}
	return defaultParams
}
