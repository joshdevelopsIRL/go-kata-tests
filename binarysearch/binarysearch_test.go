package binarysearch

import (
	"testing"

	"github.com/joshdevelopsIRL/jeer"
)

func TestBinarySearch(t *testing.T) {
    foo := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

    jeer.Test[bool](t).Actual(binarySearch(foo, 69)).Expected(true).Run("test 1")
    jeer.Test[bool](t).Actual(binarySearch(foo, 1336)).Expected(false).Run("test 2")
    jeer.Test[bool](t).Actual(binarySearch(foo, 69420)).Expected(true).Run("test 3")
    jeer.Test[bool](t).Actual(binarySearch(foo, 69421)).Expected(false).Run("test 4")
    jeer.Test[bool](t).Actual(binarySearch(foo, 1)).Expected(true).Run("test 5")
    jeer.Test[bool](t).Actual(binarySearch(foo, 0)).Expected(false).Run("test 6")
}
