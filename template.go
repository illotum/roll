package roll

import (
	"errors"
	"strings"
	"text/template"

	"github.com/illotum/roll/dice"
)

type templ struct {
	template.Template
}

// UnmarshalTOML parses string into text/template instance.
func (t *templ) UnmarshalTOML(data interface{}) error {
	tmp := template.New("")
	tmp.Funcs(templateFuncs)

	text, ok := data.(string)
	if !ok {
		return errors.New("'text' string is required in a table")
	}
	_, err := tmp.Parse(text)
	t.Template = *tmp
	return err
}

var templateFuncs = template.FuncMap{
	"random": func(t table) string {
		return t.Sample()
	},
	"pick": func(t table, n int) string {
		return t.Pick(n - 1)
	},
	"roll": func(s string) (int, error) {
		r := strings.NewReader(s)
		i, err := dice.ParseReader("", r)
		return i.(int), err
	},
}
