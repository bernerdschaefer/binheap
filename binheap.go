// The binheap package provides a min-heap implementation for storing arbitrary
// values with priority values. Unlike other heap packages, the priority is
// defined as separate from the value stored, where the priority can
// be any integer, and the smaller the number the higher the priority.
package binheap

type item struct {
	priority int
	value    interface{}
}

type list []item

type Heap struct {
	list list
}

// Allocate a new heap.
func New() *Heap {
	return &Heap{}
}

// Insert a value into the heap with the specified priority.
func (h *Heap) Insert(priority int, value interface{}) {
	item := item{priority, value}
	h.list = append(h.list, item)
	h.list.swim(len(h.list) - 1)
}

// Return the next highest priority item without removing it. Returns
// nil if no items are available.
func (h *Heap) Peek() interface{} {
	if len(h.list) == 0 {
		return nil
	}

	return h.list[0].value
}

// Remove the next highest priority item and return its value. Returns
// nil if no items are available.
// will be true, otherwise it will be false.
func (h *Heap) Pop() interface{} {
	if len(h.list) == 0 {
		return nil
	}

	// capture the top item
	value := h.list[0].value

	// remove the top item
	h.list.remove(0)

	return value
}

// Remove an item from the heap by its value.
func (h *Heap) Remove(value interface{}) {
	for index, item := range h.list {
		if item.value == value {
			h.list.remove(index)
			return
		}
	}
}

func (l *list) remove(index int) {
	list := *l

	list[index] = list[len(list)-1]
	list = list[:len(list)-1]
	list.sink(index)

	*l = list
}

func (l list) sink(index int) {
	length := len(l)
	left, right := index*2+1, index*2+2

	if left >= length {
		return
	}

	swap_index := left

	if right < length && l[right].priority < l[left].priority {
		swap_index = right
	}

	if l[swap_index].priority < l[index].priority {
		l.swap(index, swap_index)
		l.sink(swap_index)
	}
}

func (l list) swim(index int) {
	if index == 0 {
		return
	}

	parent := (index-1)/2

	if l[index].priority < l[parent].priority {
		l.swap(parent, index)
		l.swim(parent)
	}
}

func (l list) swap(a, b int) {
	l[a], l[b] = l[b], l[a]
}
