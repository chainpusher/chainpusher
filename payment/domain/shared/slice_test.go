package shared_test

import (
	"github.com/chainpusher/chainpusher/payment/domain/shared"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

type ExampleStruct struct {
	Ok bool
}

func TestSlice_Paginate(t *testing.T) {
	s := shared.Slice[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pages := s.Paginate(5)
	assert.Equal(t, 2, len(pages))
	assert.Equal(t, shared.Slice[int]{1, 2, 3, 4, 5}, pages[0])
	assert.Equal(t, shared.Slice[int]{6, 7, 8, 9, 10}, pages[1])
}

func TestSlice_Map(t *testing.T) {
	s := shared.Slice[int]{1, 2, 3, 4}
	mapped := shared.Map(s, func(i int) string {
		return strconv.Itoa(i * 2)
	})

	assert.Equal(t, shared.Slice[string]{"2", "4", "6", "8"}, mapped)

}

func TestSlice_Add(t *testing.T) {
	s := shared.Slice[*ExampleStruct]{{Ok: true}}
	s = s.Add(nil)
	assert.Equal(t, 1, len(s))
}

func TestSlice_Filter(t *testing.T) {
	s := shared.Slice[int]{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}

	filtered := s.Filter(func(i int) bool {
		return i%2 == 0
	})

	assert.Equal(t, shared.Slice[int]{0, 2, 4, 6, 8}, filtered)
}

func TestSlice_ToSlice(t *testing.T) {
	//s := shared.Slice[int]{1, 2}
	//assert.Equal(t, []int{1, 2}, s.ToSlice())
}
