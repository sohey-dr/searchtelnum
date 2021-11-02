package searchtelnum_test

import (
	"testing"
	"github.com/sohey-dr/searchtelnum"
)

func TestRunSuccess(t *testing.T) {
	t.Parallel()
	telNum, err := searchtelnum.Run("株式会社ビッグゲート", "〒905-0401")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if telNum != "0120-954-062" {
		t.Errorf("Unmatched: %v", telNum)
	}
}

func TestRunFailWithCompanyNameEmpty(t *testing.T) {
	t.Parallel()
	telNum, err := searchtelnum.Run("", "〒905-0401")
	if err == nil {
		t.Errorf("Error: %v", err)
	}

	if telNum != "" {
		t.Errorf("Error: %v", telNum)
	}
}

func TestRunFailWithPostalCodeEmpty(t *testing.T) {
	t.Parallel()
	telNum, err := searchtelnum.Run("株式会社ビッグゲート", "")
	if err == nil {
		t.Errorf("Error: %v", err)
	}

	if telNum != "" {
		t.Errorf("Error: %v", telNum)
	}
}