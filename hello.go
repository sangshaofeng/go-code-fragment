package main

import (
	"fmt"
	"unicode/utf8"
)

// 位运算
func bitOperation() {
	var x uint8 = 1<<5 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	// %08b中08表示打印至少8个字符宽度，不足的前缀部分用0填充
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)
}

// 字符串编码
// 文本字符串通常被解释为采用UTF8编码的Unicode码点（rune）序列
// 对于非ASCII字符的UTF8编码会要两个或多个字节，比如一个汉字占三个字节(byte)。
func stringCode() {
	a := "Hello"
	b := "你好"
	fmt.Println(len(a), len(b)) // 5  6
	fmt.Println(utf8.RuneCountInString(b)) // 2  通过utf8包中的方法打印出真正的unicode字符数
}

// 测试一个字符串是否是另一个字符串的前缀
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// 字符串后缀测试
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s) - len(suffix):] == suffix
}

// 包含子串测试
func Contains(s, substr  string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

// append函数
func appendTo() {
	var runes []int32
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
	fmt.Println(runes)
}

func main() {
	stringCode()
}
