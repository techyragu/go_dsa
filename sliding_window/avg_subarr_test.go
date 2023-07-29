package sliding_window_test

import (
	"log"
	"testing"

	"github.com/techyragu/go_dsa/sliding_window"
)

func TestAvgSubarr(t *testing.T) {
	input := []int{1, 3, 2, 6, -1, 4, 1, 8, 2}
	log.Print(sliding_window.FindAvgSubArray(input, 5))
}
