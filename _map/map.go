// 哈希表是一种巧妙并且实用的数据结构。它是一个无序的key/value对的集合，其中所有的key都是不同的，
// 然后通过给定的key可以在常数时间复杂度内检索、更新或删除对应的value

// 在Go语言中，一个map就是一个哈希表的引用，map类型可以写为map[K]V，其中K和V分别对应key和value
// 其中K对应的key必须是支持==比较运算符的数据类型，所以map可以通过测试key是否 相等来判断是否已经存在.

// 禁止对map元素取址的原因是map可能随着元素数量的增长而重新分配更大的内存空间，从而可能导致之前的地址无效
// Map的迭代顺序是不确定的，并且不同的哈希函数实现可能导致不同的遍历顺序。在实践中，遍历的顺序是随机的，每一次遍历的顺序都不相同。
// 这是故意的，每次都使用随机的遍历顺序可以强制要求程序不会依赖具体的哈希函数实现。

package _map

import (
	"fmt"
	"sort"
)

// 如果要按顺序遍历key/value对，我们必须显式地对key进行排序，
// 可以使用sort包的Strings函数对字符串slice进行排序。
// 在下面的第一个range循环中，我们只关心map中的key，所以我们忽略了第二个循环变量。
// 在第二个循环 中，我们只关心names中的名字，所以我们使用“_”空白标识符来忽略第一个循环变量，也就是迭代 slice时的索引。
// go语言遍历数组时，两个参数为 index(索引)，value(每一项的值)
// go语言遍历map时，两个参数为 key(键)，value(值)
func FmtSortedMap(m map[string]string) {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("%s\t%d\n", key, m[key])
	}
}


