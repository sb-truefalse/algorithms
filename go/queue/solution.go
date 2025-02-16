// Data structure: Queue
package algo

import (
	"errors"
)

/* Stack */

type Stack struct {
	items []int
}

type StackBehaviour interface {
	Pop() (int, error)
	Push(int)
	Size() int
}

func NewStack() *Stack {
	return &Stack{
		items: make([]int, 0),
	}
}

func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return -1, errors.New("empty stack")
	}

	data := s.items[s.Size()-1]
	s.items = s.items[:s.Size()-1]

	return data, nil
}

func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

func (s *Stack) Size() int {
	return len(s.items)
}

/* Queue */

type Queue struct {
	currentSize int
	pushStack   *Stack
	popStack    *Stack
}

type QueueBehaviour interface {
	Top() (int, error)
	Push(int)
	Size() int
}

func NewQueue() *Queue {
	return &Queue{
		pushStack:   NewStack(),
		popStack:    NewStack(),
		currentSize: 0,
	}
}

func (q *Queue) Top() (int, error) {
	if q.Size() == 0 {
		return -1, errors.New("empty queue")
	}

	if q.popStack.Size() == 0 {
		for data, err := q.pushStack.Pop(); err == nil; data, err = q.pushStack.Pop() {
			q.popStack.Push(data)
		}
	}

	data, err := q.popStack.Pop()
	if err != nil {
		return -1, err
	}

	q.currentSize--
	return data, nil
}

func (q *Queue) Push(data int) {
	q.pushStack.Push(data)
	q.currentSize++
}

func (q *Queue) Size() int {
	return q.currentSize
}
