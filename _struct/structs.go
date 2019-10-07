package structs

// 二叉树插入排序
type Tree struct {
	value        int
	left, right  *Tree
}

func Sort(values []int) {
	var root *Tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// 遍历二叉树，将值顺序插入到数组中
func appendValues(values []int, t *Tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

// 构造二叉树
func add(t *Tree, value int) *Tree {
	if t == nil {
		t = new(Tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	}
	if value > t.value {
		t.right = add(t.right, value)
	}
	return t
}