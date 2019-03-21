package queue

import (
	"errors"
	"sync"
)

// CircularQueue 循环队列
type CircularQueue struct {
	data         []interface{} // 容器
	front        int           // 头
	rear         int           // 尾
	max          int           // 容器容量
	sync.RWMutex               // 读写锁
}

// InitCircularQueue 初始化队列
func InitCircularQueue(lenght int) (queue *CircularQueue, err error) {
	if lenght < 1 {
		err = errors.New("长度不正确")
		return
	}
	queue = &CircularQueue{
		data:  make([]interface{}, lenght+1),
		front: 0,
		rear:  0,
		max:   lenght + 1,
	}
	return
}

// Lenght 获取当前队列长度
func (q *CircularQueue) Lenght() int {
	q.RLock()
	defer q.RUnlock()
	return (q.rear - q.front + q.max) % q.max
}

// IsFull 是否满
func (q *CircularQueue) IsFull() bool {
	q.RLock()
	defer q.RUnlock()
	return q.isFull()
}

// isFull 是否满
func (q *CircularQueue) isFull() bool {
	return ((q.rear + 1) % q.max) == q.front
}

// IsEmpty 是否空
func (q *CircularQueue) IsEmpty() bool {
	q.RLock()
	defer q.RUnlock()
	return q.isEmpty()
}

// isEmpty 是否空
func (q *CircularQueue) isEmpty() bool {
	return q.front == q.rear
}

// Enqueue 入队
func (q *CircularQueue) Enqueue(item interface{}) error {
	q.Lock()
	defer q.Unlock()
	if q.isFull() {
		return errors.New("队列已满")
	}
	q.data[q.rear] = item
	q.rear = (q.rear + 1) % q.max
	return nil
}

// Dequeue 出队
func (q *CircularQueue) Dequeue() (item interface{}, err error) {
	q.Lock()
	defer q.Unlock()
	if q.isEmpty() {
		err = errors.New("队列为空")
		return
	}
	item = q.data[q.front]
	q.front = (q.front + 1) % q.max
	return
}
