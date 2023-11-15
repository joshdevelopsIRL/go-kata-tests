package quicksort

import (
	"testing"

	"github.com/joshdevelopsIRL/jeer"
)

func TestQuicksort(t *testing.T) {
    arr := []int{9, 3, 7, 4, 69, 420, 42}
    expected := []int{3, 4, 7, 9, 42, 69, 420}
    quickSort(arr)

    jeer.Test[int](t).IsList().Actual(arr...).Expected(expected...).Run("test 1")
}
