package mylibtest

// To avoid import cycle, this package should be named {yourpackage}test.

import (
	"testing"

	"github.com/otiai10/cgosample/mylib"
)

func TestSum(t *testing.T) {
	if mylib.Sum(1, 3) != 4 {
		t.Fail()
	}
	if mylib.Sum(2, 4) != 6 {
		t.Fail()
	}
}
