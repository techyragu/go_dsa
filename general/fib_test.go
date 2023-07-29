package general_test

import (
	"testing"

	"github.com/techyragu/go_dsa/general"
)

func TestFibNumber(t *testing.T) {
	t.Log("output is ", general.GetFibNumber(3))
}

func TestFibIter(t *testing.T) {
	t.Log("output is ", general.GetFibIter(3))
}
