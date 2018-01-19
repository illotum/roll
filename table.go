package roll

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"regexp"
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
// A line number prefix "%d." is discarded if found. See `reLine`.
func (t *table) UnmarshalTOML(data interface{}) error {
	var (
		l  string
		w  int
		ok bool
	)
	switch ss := data.(type) {
	case []interface{}:
		for _, s := range ss {
			if l, ok = s.(string); !ok {
				return errors.New("table can hold only string items")
			}
			w, l = parseLine(l)
			t.ls = append(t.ls, l)
			t.ws = append(t.ws, w)
			t.total += w
		}
	case string:
		file, err := os.Open(ss)
		if err != nil {
			return fmt.Errorf("expected %q to be a file", ss)
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			w, l = parseLine(scanner.Text())
			t.ls = append(t.ls, l)
			t.ws = append(t.ws, w)
			t.total += w
		}
		if err := scanner.Err(); err != nil {
			return fmt.Errorf("unexpected error while reading %q", ss)
		}
	default:
		return fmt.Errorf("unexpected record type: %T, only path to file or list of strings are accepted", data)
	}

	if len(t.ls) == 0 {
		return fmt.Errorf("empty tables are not allowed")
	}

	return nil
}

var reLine = regexp.MustCompile(`^([0-9 ]+\. *)?([0-9]+:)?(.+)$`)

func parseLine(s string) (int, string) {
	s = strings.TrimSpace(s)
	switch len(s) {
	case 0, 1, 2:
		return 1, s
	}
	subs := reLine.FindStringSubmatch(s)
	if len(subs[2]) == 0 {
		return 1, strings.TrimSpace(subs[3])
	}

	sw := subs[2][:len(subs[2])-1]
	w, err := strconv.Atoi(sw)
	if err != nil {
		panic("error parsing priority " + sw)
	}

	return w, strings.TrimSpace(subs[3])
}
