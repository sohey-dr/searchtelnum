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
		{[]string{"帝京大学八王子", "〒192-0395"}, "0120-508-739"},
		{[]string{"", "〒192-0395"}, ""},
		{[]string{"帝京大学八王子", ""}, ""},
	}

	for _, c := range cases {
		got, _ := searchtelnum.Run(c.input[0], c.input[1])

		if got != c.want {
			t.Errorf("Run(%v) == %v, want %v", c.input, got, c.want)
		}
	}
}

func TestRunFailEmptyCompanyNamePostalCode(t *testing.T) {
	cases := []struct {
		input []string
		want  string
	}{
		{[]string{"", ""}, "company name is empty"},
		{[]string{"", "〒905-0401"}, "company name is empty"},
		{[]string{"帝京大学八王子", ""}, "postal code is empty"},
	}

	for _, c := range cases {
		_, got := searchtelnum.Run(c.input[0], c.input[1])

		if got == nil {
			t.Errorf("Run(%v) == %v, want %v", c.input, got, c.want)
		}
	}
}
