package searchtelnum_test

import (
	"testing"
	"github.com/sohey-dr/searchtelnum"
)

func TestRunSuccess(t *testing.T) {
	cases := []struct {
		input []string
		want  string
	}{
		{[]string{"株式会社ビッグゲート", "〒905-0401"}, "0120-954-062"},
		{[]string{"", "〒905-0401"}, ""},
		{[]string{"株式会社ビッグゲート", ""}, ""},
	}

	for _, c := range cases {
		got, _ := searchtelnum.Run(c.input[0], c.input[1])
		if got != c.want {
			t.Errorf("Run(%v) == %v, want %v", c.input, got, c.want)
		}
	}
}
