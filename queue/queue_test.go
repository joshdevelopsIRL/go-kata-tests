package queue

import (
	"testing"

	"github.com/joshdevelopsIRL/jeer"
)

func TestQueue(t *testing.T) {
    list := NewQueue[int]()

    list.Enqueue(5)
    list.Enqueue(7)
    list.Enqueue(9)

    result, err := list.Deque()
    jeer.Test[int](t).Actual(result).Expected(5).FailOn(err).Run("test 1")
    jeer.Test[int](t).Actual(list.Length).Expected(2).Run("test 2")

    list.Enqueue(11)
    result, err = list.Deque()
    jeer.Test[int](t).Actual(result).Expected(7).FailOn(err).Run("test 3")

    result, err = list.Deque()
    jeer.Test[int](t).Actual(result).Expected(9).FailOn(err).Run("test 4")

    result, err = list.Peek()
    jeer.Test[int](t).Actual(result).Expected(11).FailOn(err).Run("test 5")

    result, err = list.Deque()
    jeer.Test[int](t).Actual(result).Expected(11).FailOn(err).Run("test 6")

    _, res := list.Deque()
    jeer.Test[error](t).Actual(res).Expected(ERR_EMPTY).Run("test 7")
    jeer.Test[int](t).Actual(list.Length).Expected(0).Run("test 8")

    list.Enqueue(69)
    result, err = list.Peek()
    jeer.Test[int](t).Actual(result).Expected(69).FailOn(err).Run("test 9")
    jeer.Test[int](t).Actual(list.Length).Expected(1).Run("test 10")
}
