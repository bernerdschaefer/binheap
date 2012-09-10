package binheap

import "testing"

func Test_HeapInsertInOrderPopInOrder(t *testing.T) {
	heap := New()
	heap.Insert(1, "one")
	heap.Insert(2, "two")
	heap.Insert(3, "three")
	heap.Insert(4, "four")

	if value := heap.Pop(); value != "one" {
		t.Error("expected one got:", value)
	}
	if value := heap.Pop(); value != "two" {
		t.Error("expected two got:", value)
	}
	if value := heap.Pop(); value != "three" {
		t.Error("expected three got:", value)
	}
	if value := heap.Pop(); value != "four" {
		t.Error("expected four got:", value)
	}
	if value := heap.Pop(); value != nil {
		t.Error("expected value to be nil")
	}
}

func Test_HeapInsertInReverseOrderPopInOrder(t *testing.T) {
	heap := New()
	heap.Insert(4, "four")
	heap.Insert(3, "three")
	heap.Insert(2, "two")
	heap.Insert(1, "one")

	if value := heap.Pop(); value != "one" {
		t.Error("expected one got:", value)
	}
	if value := heap.Pop(); value != "two" {
		t.Error("expected two got:", value)
	}
	if value := heap.Pop(); value != "three" {
		t.Error("expected three got:", value)
	}
	if value := heap.Pop(); value != "four" {
		t.Error("expected four got:", value)
	}
	if value := heap.Pop(); value != nil {
		t.Error("expected value to be nil")
	}
}

func Test_HeapInsertInRandomOrderPopInOrder(t *testing.T) {
	heap := New()
	heap.Insert(3, "three")
	heap.Insert(4, "four")
	heap.Insert(1, "one")
	heap.Insert(2, "two")
	heap.Insert(4, "four")
	heap.Insert(1, "one")
	heap.Insert(3, "three")
	heap.Insert(2, "two")

	if value := heap.Pop(); value != "one" {
		t.Error("expected one got:", value)
	}
	if value := heap.Pop(); value != "one" {
		t.Error("expected one got:", value)
	}
	if value := heap.Pop(); value != "two" {
		t.Error("expected two got:", value)
	}
	if value := heap.Pop(); value != "two" {
		t.Error("expected two got:", value)
	}
	if value := heap.Pop(); value != "three" {
		t.Error("expected three got:", value)
	}
	if value := heap.Pop(); value != "three" {
		t.Error("expected three got:", value)
	}
	if value := heap.Pop(); value != "four" {
		t.Error("expected four got:", value)
	}
	if value := heap.Pop(); value != "four" {
		t.Error("expected four got:", value)
	}
	if value := heap.Pop(); value != nil {
		t.Error("expected value to be nil")
	}
}

func Test_HeapRemovePreservesOrder(t *testing.T) {
	heap := New()
	heap.Insert(1, "one")
	heap.Insert(10, "ten")
	heap.Insert(10, "ten")
	heap.Insert(2, "two")
	heap.Insert(3, "three")
	heap.Insert(4, "four")
	heap.Insert(5, "five")

	heap.Remove("two")

	if value := heap.Pop(); value != "one" {
		t.Error("expected one got:", value)
	}
	if value := heap.Pop(); value != "three" {
		t.Error("expected three got:", value)
	}
	if value := heap.Pop(); value != "four" {
		t.Error("expected four got:", value)
	}
	if value := heap.Pop(); value != "five" {
		t.Error("expected four got:", value)
	}
	if value := heap.Pop(); value != "ten" {
		t.Error("expected four got:", value)
	}
	if value := heap.Pop(); value != "ten" {
		t.Error("expected four got:", value)
	}
	if value := heap.Pop(); value != nil {
		t.Error("expected value to be nil")
	}
}

func Test_HeapPeek(t *testing.T) {
	heap := New()
	heap.Insert(20, "twenty")
	heap.Insert(5, "five")
	heap.Insert(10, "ten")
	heap.Insert(1, "one")

	if value := heap.Peek(); value != "one" {
		t.Error("expected one got:", value)
	}

	heap.Pop()
	if value := heap.Peek(); value != "five" {
		t.Error("expected five got:", value)
	}

	heap.Pop()
	if value := heap.Peek(); value != "ten" {
		t.Error("expected ten got:", value)
	}

	heap.Pop()
	if value := heap.Peek(); value != "twenty" {
		t.Error("expected twenty got:", value)
	}

	heap.Pop()
	if value := heap.Peek(); value != nil {
		t.Error("expected value to be nil")
	}
}
