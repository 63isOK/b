package b_test

import (
	"testing"

	. "github.com/63isOK/b"
)

func TestGetSum(t *testing.T) {
	got := GetSum()
	if got != 3 {
		t.Fatalf("want: 3, got: %d", got)
	}
}
