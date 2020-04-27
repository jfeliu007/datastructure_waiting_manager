package datastructures

import (
	"fmt"
	"github.com/pkg/errors"
)

type node struct {
	next *node
	previous *node
	value interface{}
}

type Queue struct {
	head *node
	tail *node
	len  int
}

func (q *Queue) GetAndRemove() (interface{}, error) {
	if q.head == nil {
		return nil, fmt.Errorf("empty queue")
	}
	q.len -= 1
	result := q.head.value
	q.head = q.head.previous
	if q.head != nil {
		q.head.next = nil
	} else {
		q.tail = nil
	}
	return result, nil
}

func (q *Queue) Add(value interface{}) error {
	q.len += 1
	newNode := &node{
		value:value,
		next:q.tail,
	}
	if q.head == nil{
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.previous = newNode
		q.tail = newNode
	}
	return nil
}

func (q *Queue) IsEmpty() bool {
	return q.len == 0
}

func (q *Queue) Clear() {
	q.head = nil
	q.tail = nil
	q.len = 0
}

func (q *Queue) Peek() (interface{}, error) {
	if q.head == nil {
		return nil, errors.New("empty queue")
	} else {
		return q.head.value, nil
	}
}
func (q *Queue) Size() int {
	return q.len
}

func NewQueue() *Queue {
	return &Queue{
		head: nil,
		tail: nil,
		len:  0,
	}
}
