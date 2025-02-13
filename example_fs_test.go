package extemplate_test

import (
	"embed"
	"html/template"
	"os"
	"strings"

	"github.com/dannyvankooten/extemplate"
)

//go:embed examples/*.tmpl examples/**/*.tmpl
var FS embed.FS

func ExampleFsExtemplate_ParseDir() {
	xt := extemplate.New().Funcs(template.FuncMap{
		"tolower": strings.ToLower,
	})

	err := xt.ParseDir("examples", []string{".tmpl"}, &FS)
	if err != nil {
		panic(err)
	}

	err = xt.ExecuteTemplate(os.Stdout, "child.tmpl", nil)
	if err != nil {
		panic(err)
	}
	/* Output: Hello from child.tmpl
	Hello from partials/question.tmpl */
}
