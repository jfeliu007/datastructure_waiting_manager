package datastructures

type DataStructure interface {
	GetAndRemove()(interface{}, error)
	Add(interface{}) error
	IsEmpty() bool
	Clear()
	Peek()(interface{}, error)
	Size() int
}