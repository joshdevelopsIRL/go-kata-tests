package stack

import "errors"

var ERR_EMPTY = errors.New("node stack is empty")

type Node[T any] struct {

}

type Stack[T any] struct {
    Length int
}

func NewStack[T any]() *Stack[T] {

}

func (q *Stack[T]) Push(item T) {

}

func (q *Stack[T]) Pop() (T, error) {

}

func (q *Stack[T]) Peek() (T, error) {

}
