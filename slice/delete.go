package slice

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
