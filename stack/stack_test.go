package stack

import (
	"testing"

	"github.com/joshdevelopsIRL/jeer"
)

func TestStack(t *testing.T) {
    list := NewStack[int]()

    list.Push(5)
    list.Push(7)
    list.Push(9)

    result, err := list.Pop()
    jeer.Test[int](t).Actual(result).Expected(9).FailOn(err).Run("test 1")
    jeer.Test[int](t).Actual(list.Length).Expected(2).Run("test 2")

    list.Push(11)
    result, err = list.Pop()
    jeer.Test[int](t).Actual(result).Expected(11).FailOn(err).Run("test 3")

    result, err = list.Pop()
    jeer.Test[int](t).Actual(result).Expected(7).FailOn(err).Run("test 4")

    result, err = list.Peek()
    jeer.Test[int](t).Actual(result).Expected(5).FailOn(err).Run("test 5")

    result, err = list.Pop()
    jeer.Test[int](t).Actual(result).Expected(5).FailOn(err).Run("test 6")

    _, res := list.Pop()
    jeer.Test[error](t).Actual(res).Expected(ERR_EMPTY).Run("test 7")
    jeer.Test[int](t).Actual(list.Length).Expected(0).Run("test 8")

    list.Push(69)
    result, err = list.Peek()
    jeer.Test[int](t).Actual(result).Expected(69).FailOn(err).Run("test 9")
    jeer.Test[int](t).Actual(list.Length).Expected(1).Run("test 10")
}
