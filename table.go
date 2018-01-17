package roll

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
)

// table acts as a weighted string list.
type table struct {
	ls    []string
	ws    []int
	total int
}

// List returns a copy of underlying string list.
func (t table) List() []string {
	return t.ls[:]
}

// Total returns weighted table length.
func (t table) Total() int {
	return t.total
}

// Pick returns string at the given weighted table pos.
func (t table) Pick(w int) string {
	if w > t.Total() {
		w = t.Total()
	}
	return t.ls[t.findIx(w)]
}

// Sample returns a weighted random choice out of the list.
func (t table) Sample() string {
	r := rand.Intn(t.total)
	return t.Pick(r)
}

// findIx translates weights into list indices for a given table.
func (t table) findIx(w int) int {
	for i := range t.ws {
		w -= t.ws[i]
		if w < 0 {
			return i
		}
	}
	panic("weight out of table bounds")
}

// UnmarshalTOML parses string list into a weighted table.
//
// Strings in "%d:%s" form will have integer weight parsed out and trimmed,
// all other forms are taken as-is with a weight of 1.
//
// TODO: fix unnecessary allocations
func (t *table) UnmarshalTOML(data interface{}) error {
	var (
		l   string
		w   int
		ok  bool
		err error
		ss  []interface{}
	)
	if ss, ok = data.([]interface{}); !ok {
		return errors.New("table should be a list of strings")
	}
	for _, s := range ss {
		if l, ok = s.(string); !ok {
			return errors.New("table can hold only string items")
		}
		i := strings.Index(l, ":")
		if i == -1 {
			t.ls = append(t.ls, l)
			t.ws = append(t.ws, 1)
			t.total++
			continue
		}

		w, err = strconv.Atoi(l[:i])
		if err != nil {
			t.ls = append(t.ls, l)
			t.ws = append(t.ws, 1)
			t.total++
			continue
		}

		l = strings.Trim(l[i:], "\t\n\r :")
		t.ls = append(t.ls, l)
		t.ws = append(t.ws, w)

		t.total += w
	}

	return nil
}
