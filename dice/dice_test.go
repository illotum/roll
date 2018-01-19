package dice_test

import (
	"math/rand"
	"testing"

	"github.com/illotum/roll/dice"
)

type diceCase struct {
	str string
	exp int
}

func TestParser(t *testing.T) {
	var cases = []diceCase{
		{"1", 1},
		{"1+1", 2},
		{"1*1", 1},
		{"1/2", 0},
		{"2/2", 1},
		{"2+(2-2)", 2},
		{"2*(2-2)", 0},
		{"2*2-2", 2},
		{"6-(2*2)", 2},
		{"1d1", 1},
		{"(1+1)d1", 2},
		{"(1+1)d(1*1)", 2},
		{"0d100", 0},
		{"1d6", 1},
		{"10*1d6", 10},
	}
	rand.Seed(1)

	for _, c := range cases {
		t.Run(c.str, func(t *testing.T) {
			rand.Seed(0)
			i, err := dice.Parse("", []byte(c.str))
			if err != nil {
				t.Fatalf("unexpected error whle parsing %q: %s", c.str, err)
			}
			res := i.(int)
			if res != c.exp {
				t.Errorf("expected %d from parsing %s, got %d", c.exp, c.str, res)
			}
		})
	}
}
