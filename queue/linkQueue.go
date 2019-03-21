package queue

import (
	"errors"

	"github.com/zboyco/goutlis/list"
)

// LinkQueue 链表队列
type LinkQueue struct {
	list.SinglyList
}

// Enqueue 入队
func (q *LinkQueue) Enqueue(item interface{}) error {
	return q.Append(item)
}

// Dequeue 出队
func (q *LinkQueue) Dequeue() (item interface{}, err error) {
	item, err = q.Remove(0)
	if err != nil {
		err = errors.New("队列为空")
	}
	return
}
