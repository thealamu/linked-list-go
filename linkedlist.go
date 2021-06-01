package linkedlists

type node struct {
	data interface{}
	next *node
}

type LinkedLists struct {
	head *node
}

func New() *LinkedLists {
	return &LinkedLists{}
}

// PushFront adds an item to the front of the list
func (l *LinkedLists) PushFront(item interface{}) {
	if l.head == nil {
		l.head = &node{data: item}
	}

	newHead := &node{data: item}
	newHead.next = l.head
	l.head = newHead
}
