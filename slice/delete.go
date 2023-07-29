package slice

// Slicer 泛型切片类型
type Slicer[T any] []T

// RemoveAt 基础版本
func RemoveAt(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

// RemoveFast 高性能版本，并增加了索引检查
func RemoveFast(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice // 索引无效，直接返回原切片
	}

	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}

// RemoveGeneric 泛型版本
func RemoveGeneric[T any](s Slicer[T], index int) Slicer[T] {
	if index < 0 || index >= len(s) {
		return s // 索引无效，直接返回原切片
	}

	// 将 Slicer[T] 转为其基础的切片类型进行操作
	baseSlice := []T(s)
	copy(baseSlice[index:], baseSlice[index+1:])
	return baseSlice[:len(baseSlice)-1]
}

// TrimRemove 实现缩容机制
const (
	thresholdRatio    = 0.25
	maxReductionCount = 3
	bigSliceCapacity  = 10000
)

var reductionCount = 0

func TrimRemove[T any](s Slicer[T], index int) Slicer[T] {
	if index < 0 || index >= len(s) {
		return s // 索引无效，直接返回原切片
	}

	baseSlice := []T(s)
	copy(baseSlice[index:], baseSlice[index+1:])
	result := baseSlice[:len(baseSlice)-1]

	if cap(result) > bigSliceCapacity && float64(len(result))/float64(cap(result)) < thresholdRatio {
		reductionCount++
		if reductionCount >= maxReductionCount {
			newSlice := make([]T, len(result), len(result)*2)
			copy(newSlice, result)
			reductionCount = 0
			return newSlice
		}
	} else {
		reductionCount = 0
	}

	return result
}
