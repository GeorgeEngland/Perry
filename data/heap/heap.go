package heap

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// The Interface type describes the requirements
// for a type using the routines in this package.
// Any type that implements it may be used as a
// min-heap with the following invariants (established after
// [Init] has been called or if the data is empty or sorted):
//
//	!h.Less(j, i) for 0 <= i < h.Len() and 2*i+1 <= j <= 2*i+2 and j < h.Len()
//
// Note that [Push] and [Pop] in this interface are for package heap's
// implementation to call. To add and remove things from the heap,
// use [heap.Push] and [heap.Pop].
type Interface[T constraints.Ordered] interface {
	sort.Interface
	push(x T) // add x as element Len()
	pop() T   // remove and return element Len() - 1.
}

// Init establishes the heap invariants required by the other routines in this package.
// Init is idempotent with respect to the heap invariants
// and may be called whenever the heap invariants may have been invalidated.
// The complexity is O(n) where n = h.Len().
func Init[T constraints.Ordered](h Interface[T]) {
	// heapify
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func Push[T constraints.Ordered](h Interface[T], x T) {
	h.push(x)
	up(h, h.Len()-1)
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to [Remove](h, 0).
func Pop[T constraints.Ordered](h Interface[T]) T {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.pop()
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func Remove[T constraints.Ordered](h Interface[T], i int) T {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)
		if !down(h, i, n) {
			up(h, i)
		}
	}
	return h.pop()
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling [Remove](h, i) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
func Fix[T constraints.Ordered](h Interface[T], i int) {
	if !down(h, i, h.Len()) {
		up(h, i)
	}
}

func up[T constraints.Ordered](h Interface[T], j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

func down[T constraints.Ordered](h Interface[T], i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}

// we need to define a custom type instead of using the raw integer slice
// since we need to define methods on the type to implement the heap interface
type Heap[T constraints.Ordered] []T

// Len is the number of elements in the collection.
func (h Heap[T]) Len() int {
	return len(h)
}

// Less reports whether the element with index i
// must sort before the element with index j.
// This will determine whether the heap is a min heap or a max heap
func (h Heap[T]) Less(i int, j int) bool {
	return h[i] < h[j]
}

// Swap swaps the elements with indexes i and j.
func (h Heap[T]) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push and Pop are used to append and remove the last element of the slice
func (h *Heap[T]) push(x T) {
	*h = append(*h, x)
}

func (h *Heap[T]) pop() T {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
