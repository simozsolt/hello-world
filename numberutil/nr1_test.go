package numberutil

import "testing"

func TestNrDividerRatioRest(t *testing.T) {
	cases := []struct{
		nr int
		divider int
		ratio int
		rest int
		something string
	}{
		{100, 10, 10, 0, "1"},
		{122, 10, 12, 2, "2"},
	}
	for _, c := range cases {
		gotRatio, gotRest := NrDividerRatioRest(c.nr, c.divider)
		if gotRatio != c.ratio || gotRest != c.rest {
			t.Errorf("Something wron in case: %q", c.something)
		}
	}
}
