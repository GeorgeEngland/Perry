package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	nums := &Heap[int]{3, 1, 4, 5, 1, 1, 2, 6}
	// The `Init` function reorders the numbers into a heap[T]
	Init(nums)
	// The slice is now reordered to conform to the heap property
	prev_num := 0
	i := 0
	length := 8
	for nums.Len() > 0 {
		new_num := Pop(nums)
		assert.LessOrEqual(t, prev_num, new_num)
		prev_num = new_num
		i++
	}
	assert.Equal(t, length, i)
}
