package datastructures

import (
"fmt"
"github.com/pkg/errors"
)


type Stack struct {
	head *node
	len  int
}

func (q *Stack) GetAndRemove() (interface{}, error) {
	if q.head == nil {
		return nil, fmt.Errorf("empty queue")
	}
	q.len -= 1
	result := q.head.value
	q.head = q.head.next
	return result, nil
}

func (q *Stack) Add(value interface{}) error {
	q.len += 1
	newNode := &node{
		value:value,
		next:q.head,
	}
	if q.head == nil{
		q.head = newNode
	} else {
		q.head.next = newNode
		q.head = newNode
	}
	return nil
}

func (q *Stack) IsEmpty() bool {
	return q.len == 0
}

func (q *Stack) Clear() {
	q.head = nil
	q.len = 0
}

func (q *Stack) Peek() (interface{}, error) {
	if q.head == nil {
		return nil, errors.New("empty queue")
	} else {
		return q.head.value, nil
	}
}
func (q *Stack) Size() int {
	return q.len
}

func NewStack() *Stack {
	return &Stack{
		head: nil,
		len:  0,
	}
}
