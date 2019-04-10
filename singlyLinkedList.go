package goutlis

import (
	"errors"
	"sync"
)

type singlyNode struct {
	value interface{}
	next  *singlyNode
}

// SinglyList 单向链表
type SinglyList struct {
	lenght int
	front  *singlyNode
	rear   *singlyNode
	sync.RWMutex
}

// Lenght 当前长度
func (list *SinglyList) Lenght() int {
	return list.lenght
}

// Append 追加数据
func (list *SinglyList) Append(item interface{}) error {
	if item == nil {
		return errors.New("数据不可为空")
	}
	list.Lock()
	defer list.Unlock()
	newNode := &singlyNode{value: item}

	if list.front == nil {
		list.front = newNode
	}
	if list.rear != nil {
		list.rear.next = newNode
	}
	list.rear = newNode
	list.lenght++
	return nil
}

// Get 查找元素
func (list *SinglyList) Get(index int) (item interface{}, err error) {
	list.RLock()
	defer list.RUnlock()
	if index < 0 || list.lenght <= index {
		err = errors.New("索引超出界限")
		return
	}
	item = list.findNode(index).value
	return
}

// InsertAt 插入数据
func (list *SinglyList) InsertAt(index int, item interface{}) error {
	list.Lock()
	defer list.Unlock()
	if list.lenght == 0 {
		return errors.New("空链表无法插入")
	}
	if index < 0 || list.lenght <= index {
		return errors.New("索引超出界限")
	}
	newNode := &singlyNode{value: item}
	if index == 0 {
		newNode.next = list.front
		list.front = newNode
	} else {
		preNode := list.findNode(index - 1)
		newNode.next = preNode.next
		preNode.next = newNode
	}
	list.lenght++
	return nil
}

// Remove 移除元素
func (list *SinglyList) Remove(index int) (interface{}, error) {
	list.Lock()
	defer list.Unlock()
	if index < 0 || list.lenght <= index {
		return nil, errors.New("索引超出界限")
	}
	var curNode *singlyNode
	if index == 0 {
		curNode = list.front
		list.front = list.front.next
		if list.lenght == 1 {
			list.rear = nil
		}
	} else {
		preNode := list.findNode(index - 1)
		curNode = preNode.next
		preNode.next = curNode.next
		if index == (list.lenght - 1) {
			list.rear = preNode
		}
	}
	list.lenght--
	return curNode.value, nil
}

// Clear 清空链表
func (list *SinglyList) Clear() {
	list.Lock()
	defer list.Unlock()
	list.front = nil
	list.rear = nil
	list.lenght = 0
	return
}

func (list *SinglyList) findNode(index int) (n *singlyNode) {
	n = list.front
	for i := 0; i < index; i++ {
		n = n.next
	}
	return
}
