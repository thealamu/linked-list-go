package linkedlists

import "testing"

func TestPushFront(t *testing.T) {
	l := New()

	v1 := 1
	v2 := "Hello"
	v3 := "World"

	l.PushFront(v1)
	l.PushFront(v2)
	l.PushFront(v3)

	if l.head.data != v3 {
		t.Errorf("PushFront failed, expected '%+v' to be in front, got '%+v'", v3, l.head.data)
	}
}

func TestSize(t *testing.T) {
	l := New()

	v1 := 1
	v2 := "Hello"
	v3 := "World"

	l.PushFront(v1)
	l.PushFront(v2)
	l.PushFront(v3)

	got := l.Size()

	if got != 3 {
		t.Errorf("expected size %d, got %d", 3, got)
	}
}
