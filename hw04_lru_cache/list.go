package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len   int
	front *ListItem
	back  *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{Value: v, Next: l.front}

	if l.front != nil {
		l.front.Prev = item
	}

	if l.back == nil {
		l.back = item
	}

	l.front = item
	l.len++

	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := &ListItem{Value: v, Prev: l.back}

	if l.back != nil {
		l.back.Next = item
	}

	if l.front == nil {
		l.front = item
	}

	l.back = item
	l.len++

	return l.back
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.front = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.back = i.Prev
	}

	l.len--
	i = nil
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)

	i.Prev = nil
	i.Next = l.front

	if l.front != nil {
		l.front.Prev = i
	}

	if l.back == nil {
		l.back = i
	}

	l.front = i
	l.len++
}

func NewList() List {
	return new(list)
}
