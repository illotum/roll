package main

import (
	"strings"
	"text/template"

	"github.com/illotum/roll/dice"
	"github.com/illotum/roll/table"
)

func parseTemplate(text string) (*template.Template, error) {
	t := template.New("")
	t.Funcs(template.FuncMap{
		"random": func(t table.Table) string {
			return t.Sample()
		},
		"pick": func(t table.Table, n int) string {
			return t.Pick(n - 1)
		},
		"roll": func(s string) (int, error) {
			r := strings.NewReader(s)
			i, err := dice.ParseReader("", r)
			return i.(int), err
		},
	})
	return t.Parse(text)
}
