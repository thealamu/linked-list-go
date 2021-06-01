package linkedlists

type LinkedLists struct {
	head *node
}

type node struct {
	data interface{}
	next *node
}

func New() *LinkedLists {
	return &LinkedLists{}
}

// ValueAt returns the value of the item at index (starting at 0 for first)
func (l *LinkedLists) ValueAt(index int) interface{} {
	var data interface{}

	i := 0
	l.Walk(func(n *node) bool {
		if i == index {
			data = n.data
			return true
		}
		i++
		return false
	})

	return data
}

// PushFront adds an item to the front of the list
func (l *LinkedLists) PushFront(item interface{}) {
	if l.head == nil {
		l.head = &node{data: item}
		return
	}

	newHead := &node{data: item}
	newHead.next = l.head
	l.head = newHead
}

// PopFront removes front item and returns its value
func (l *LinkedLists) PopFront() interface{} {
	if l.head == nil {
		return nil
	}
	tmp := l.head
	l.head = l.head.next
	return tmp.data
}

func (l *LinkedLists) PushBack(item interface{}) {
	newNode := &node{data: item}
	if l.head == nil {
		l.head = newNode
		return
	}

	l.Walk(func(n *node) bool {
		if n.next == nil {
			// last node
			n.next = newNode
			return true
		}
		return false
	})
}

// Size returns the length of the linked lists
func (l *LinkedLists) Size() int {
	// walk the list to report the size
	i := 0
	l.Walk(func(*node) bool {
		i++
		return false
	})
	return i
}

// Returns true if list is empty
func (l *LinkedLists) Empty() bool {
	return l.Size() == 0
}

// Function to call for each iteration when walking
type iterFunc func(*node) bool

// Walk walks the linked lists, calling the iterFunc each iteration.
// Signify stopping the walk by returning true from the iterFunc
func (l *LinkedLists) Walk(f iterFunc) {
	curr := l.head
	for curr != nil {
		if f(curr) {
			break
		}
		curr = curr.next
	}
}
