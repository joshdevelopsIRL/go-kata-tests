package queue

import "errors"

var ERR_EMPTY = errors.New("node queue is empty")

type Node[T any] struct {

}

type Queue[T any] struct {
    Length int
}

func NewQueue[T any]() *Queue[T] {

}

func (q *Queue[T]) Enqueue(item T) {

}

func (q *Queue[T]) Deque() (T, error) {

}

func (q *Queue[T]) Peek() (T, error) {

}
