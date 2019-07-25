// slice的底层是引用的数组对象
// 一个slice由三部分组成：指针、长度、和容量
// 指针指向第一个slice元素对应的底层数组元素的地址，要注意的是slice的第一个元素并不一 定就是数组的第一个元素
// 长度对应slice中元素的数目；长度不能超过容量，
// 容量一般是从slice的开始位置到底层数据的结尾位置。内置的len和cap函数分别返回slice的长度和容量
// 多个slice之间可以共享底层的数据，并且引用的数组部分区间可能重叠

// slice的切片操作s[i:j]，其中0 ≤ i ≤ j ≤ cap(s)，用于创建一个新的slice，
// 引用s的从第i个元素开始到第j-1个元素的子序列，新的slice将只有j-i个元素
// 如果i位置的索引被省略的话将使用0代替，如 果j位置的索引被省略的话将使用len(s)代替

package slice

// 因为slice值包含指向第一个slice元素的指针，因此向函数传递slice将允许在函数内部修改底层数组的 元素。
// 换句话说，复制一个slice只是对底层的数组创建了一个新的slice别名。
// 下面的 reverse函数在原内存空间将[]int类型的slice反转，而且它可以用于任意长度的slice
func ReverseSlice(s []int) {
	// 直接修改了s
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 使用指针代替slice，完成上面的方法
func ReverseSlice2(s *[]byte) {
	i, j := 0, len(*s) - 1
	for i < j {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
		i += 1
		j -= 1
	}
}

// 移除slice中间的某一位，slice目标slice, i要移除元素的位置
func Remove(slice []int, i int) []int {
	// copy(destSlice, srcSlice []T) int 返回实际发生复制的元素
	// 将第二个slice里的元素拷贝到第一个slice里，拷贝的长度为两个slice中长度较小的长度值
	// 这里相当于每一项往前一位，最后返回除最后一位的slice
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

