package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/* Stack */
func TestNewStack(t *testing.T) {
	assert.IsType(t, NewStack(), &Stack{})
}

func TestPushForStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)

	assert.Equal(t, stack.items[0], 1)
	assert.Equal(t, stack.items[1], 2)
}

func TestPopForStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)

	data, err := stack.Pop()
	if assert.Nil(t, err) {
		assert.Equal(t, data, 2)
	}
}

func TestSizeForStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)

	assert.Equal(t, stack.Size(), 2)
}

/* Queue */

func TestNewQueue(t *testing.T) {
	assert.IsType(t, NewQueue(), &Queue{})
}

func TestPushForQueue(t *testing.T) {
	queue := NewQueue()
	queue.Push(1)
	queue.Push(2)

	assert.Equal(t, queue.pushStack.items[0], 1)
	assert.Equal(t, queue.pushStack.items[1], 2)
}

func TestPopForQueue(t *testing.T) {
	queue := NewQueue()
	queue.Push(1)
	queue.Push(2)

	data, err := queue.Top()
	if assert.Nil(t, err) {
		assert.Equal(t, data, 1)
	}
}

func TestSizeForQueue(t *testing.T) {
	queue := NewQueue()
	queue.Push(1)
	queue.Push(2)

	assert.Equal(t, queue.Size(), 2)
}
