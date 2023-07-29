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
