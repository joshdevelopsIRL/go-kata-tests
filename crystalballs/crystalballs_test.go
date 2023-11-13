package crystalballs

import (
	"testing"

	"github.com/joshdevelopsIRL/jeer"
)

func TestCrystalBalls(t *testing.T) {
    foo := make([]bool, 10000)
    idx := 777
    for i := idx; i < len(foo) - idx; i++ {
        foo[i] = true
    }

    jeer.Test[int](t).Actual(crystalBalls(foo)).Expected(777).Run("test 1")

    foox := make([]bool, 2000)
    jeer.Test[int](t).Actual(crystalBalls(foox)).Expected(-1).Run("test 2")
    
    fooz := make([]bool, 625)
    for i := range fooz {
        fooz[i] = true
    }
    jeer.Test[int](t).Actual(crystalBalls(fooz)).Expected(0).Run("test 3")
}
