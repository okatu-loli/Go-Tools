package slice_test

import (
	"github.com/okatu-loli/Go-Tools/slice"
	"reflect"
	"testing"
)

func TestRemoveAt(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		index    int
		expected []int
	}{
		{
			name:     "正常情况",
			slice:    []int{1, 2, 3, 4, 5},
			index:    2,
			expected: []int{1, 2, 4, 5},
		},
		{
			name:     "删除第一个元素",
			slice:    []int{1, 2, 3, 4, 5},
			index:    0,
			expected: []int{2, 3, 4, 5},
		},
		{
			name:     "删除最后一个元素",
			slice:    []int{1, 2, 3, 4, 5},
			index:    4,
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slice.RemoveAt(tt.slice, tt.index)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("期望得到 %v, 但是得到 %v", tt.expected, got)
			}
		})
	}
}
