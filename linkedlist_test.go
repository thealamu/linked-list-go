package linkedlists

import "testing"

func TestValueAt(t *testing.T) {
	l := New()

	v1 := 1
	v2 := "Hello"
	v3 := "World"

	l.PushFront(v1)
	l.PushFront(v2)
	l.PushFront(v3)

	got := l.ValueAt(1)
	if got != v2 {
		t.Errorf("expected '%+v' at index 1, got %+v", v2, got)
	}
}

func TestPopFront(t *testing.T) {
	l := New()

	v1 := 1
	v2 := "Hello"
	v3 := "World"

	l.PushFront(v1)
	l.PushFront(v2)
	l.PushFront(v3)

	got := l.PopFront()
	if got != v3 {
		t.Errorf("expected '%+v', got '%+v'", v3, got)
	}

	got = l.Size()
	if got != 2 {
		t.Errorf("expected size after PopFront to be %d, got %d", 2, got)
	}
}

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

func TestEmpty(t *testing.T) {
	l := New()

	if !l.Empty() {
		t.Errorf("expected list to be empty")
	}
}
