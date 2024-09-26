package queue

import (
	"errors"
	"sync"
)

type Queue[T any] struct {
	items []T
	mutex sync.Mutex
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: make([]T, 0)}
}

// Size returns the current number of items in the queue.
func (q *Queue[T]) Size() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	return len(q.items)
}

// IsEmpty returns true if the queue has no items.
func (q *Queue[T]) IsEmpty() bool {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	return len(q.items) == 0
}

// Enqueue adds an item to the end of the queue.
func (q *Queue[T]) Enqueue(item T) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	var zero T
	if len(q.items) == 0 {
		return zero, errors.New("queue is empty")
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}
