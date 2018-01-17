package table

import (
	"strconv"
	"strings"
)

// New transforms string list into a weighted Table.
//
// Strings in a "%d:%s" form will have integer weight parsed out and trimmed,
// all other forms are taken as-is with a weight of 1.
func New(ss []string) Table {
	var (
		ls       []string
		l        string
		ws       []int
		w, total int
		err      error
	)
	for _, s := range ss {
		i := strings.Index(s, ":")
		if i == -1 {
			ls = append(ls, s)
			ws = append(ws, 1)
			total++
			continue
		}

		w, err = strconv.Atoi(s[:i])
		if err != nil {
			ls = append(ls, s)
			ws = append(ws, 1)
			total++
			continue
		}

		l = strings.Trim(s[i:], "\t\n\r :")
		ls = append(ls, l)
		ws = append(ws, w)

		total += w
	}

	return Table{
		ls:    ls,
		ws:    ws,
		total: total,
	}
}
