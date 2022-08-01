package fibonacci

import "testing"

var cases = []struct {
	have, want int
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{12, 144},
	{19, 4181},
}

func TestCalculate(t *testing.T) {
	for _, c := range cases {
		got := Calculate(c.have)
		if got != c.want {
			t.Errorf("Calculate(%d) = %d; want %d", c.have, got, c.want)
		}
	}
}
