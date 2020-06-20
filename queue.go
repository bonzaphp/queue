//
// @Author: bonzaphp@gmail.com
// @Date: 2020/6/15 2:34 下午
//
package queue

import "errors"

//定义队列接口
type IQueue interface {
	Enqueue(i *Item)
	Dequeue() (interface{}, error)
	IsEmpty() bool
	Size() int
}

//定义队列元素
type Item struct {
	Value interface{}
	prev  *Item
	next  *Item
}

//定义队列结构
type Queue struct {
	tail *Item
	head *Item
	size int
}

// 创建一个新队列
func New() *Queue {
	return new(Queue)
}

// 入队
func (q *Queue) Enqueue(i *Item) {
	if q.tail == nil {
		q.head = i
		q.tail = i
	} else {
		q.tail.next = i
		q.tail = i
	}
	q.size++
}

// 出队
func (q *Queue) Dequeue() (interface{}, error) {
	var (
		firstNode *Item
		value     interface{}
	)
	if q.head == nil {
		return false, errors.New("这是一个空队列")
	}
	firstNode = q.head
	//这里必须要检查这一部分，因为入队的时候，是根据尾部元素是否为空来添加新元素的
	if firstNode.next == nil {
		q.tail = nil
	}
	q.head = firstNode.next
	firstNode.next = nil
	value = firstNode.Value
	firstNode.Value = nil
	q.size--
	firstNode = nil
	return value, nil
}

//是否为空
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// 队列大小
func (q *Queue) Size() int {
	return q.size
}
