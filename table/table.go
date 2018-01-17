package table

import "math/rand"

// Table acts as a weighted string list.
type Table struct {
	ls    []string
	ws    []int
	total int
}

// List returns a copy of underlying string list.
func (t Table) List() []string {
	return t.ls[:]
}

// Total returns weighted table length.
func (t Table) Total() int {
	return t.total
}

// Pick returns string at the given weighted table pos.
func (t Table) Pick(w int) string {
	if w > t.Total() {
		w = t.Total()
	}
	return t.ls[t.findIx(w)]
}

// Sample returns a weighted random choice out of the list.
func (t Table) Sample() string {
	r := rand.Intn(t.total)
	return t.Pick(r)
}

// findIx translates weights into list indices for a given table.
func (t Table) findIx(w int) int {
	for i := range t.ws {
		w -= t.ws[i]
		if w < 0 {
			return i
		}
	}
	panic("weight out of table bounds")
}
