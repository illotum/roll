package roll

import (
	"errors"

	"github.com/BurntSushi/toml"
)

// Record describes one template and corresponding weighted tables.
type Record struct {
	Template templ            `toml:"template"`
	Tables   map[string]table `toml:"tables"`
}

// ReadFile parses template and tables out of a TOML file.
func ReadFile(name string) (Record, error) {
	var f Record
	md, err := toml.DecodeFile(name, &f)
	if err != nil {
		return f, err
	}

	if !md.IsDefined("tables") {
		return f, errors.New("'text' string is required in a table")

	}

	return f, nil
}
