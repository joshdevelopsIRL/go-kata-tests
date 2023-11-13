package bubblesort

import (
	"testing"

	"github.com/joshdevelopsIRL/jeer"
)

func TestBubbleSort(t *testing.T) {
    list := []int{9, 3, 7, 4, 69, 420, 42}
    
    bubbleSort(list)
    jeer.Test[int](t).IsList().Actual(list...).Expected(3, 4, 7, 9, 42, 69, 420).Run("test 1")
}
