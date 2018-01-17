package roll

import (
	"errors"

	"github.com/BurntSushi/toml"
)

// Record describes one template and corresponding weighted tables.
type Record struct {
	Text templ            `toml:"text"`
	Data map[string]table `toml:"data"`
}

// ReadFile parses template and tables out of a TOML file.
func ReadFile(name string) (Record, error) {
	var f Record
	md, err := toml.DecodeFile(name, &f)
	if err != nil {
		return f, err
	}

	if !md.IsDefined("text") {
		return f, errors.New("'text' string is required in a table")

	}

	return f, nil
}
