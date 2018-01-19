package roll

import "testing"

type lineCase struct {
	name, s, l string
	w          int
}

func TestLineParser(t *testing.T) {
	var cases = []lineCase{
		{"empty", "", "", 1},
		{"sentence", "Test.", "Test.", 1},
		{"numeric", "42.", "42.", 1},
		{"numbered", "2. Test.", "Test.", 1},
		{"weighted", "2: Test.", "Test.", 2},
		{"numbered-weighted", "2. 3: Test.", "Test.", 3},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			w, l := parseLine(c.s)
			if w != c.w {
				t.Errorf("Got weight %d from parsing %q, but expected %d.", w, c.s, c.w)
			}
			if l != c.l {
				t.Errorf("Got line %q from parsing %q, but expected %q.", l, c.s, c.l)
			}
		})
	}

}
