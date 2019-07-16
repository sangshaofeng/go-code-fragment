package string

import (
	"fmt"
	"unicode/utf8"
)

// golang中string底层是通过byte数组实现的。中文字符在unicode下占2个字节，
// 在utf-8编码下占3个字节， golang默认编码正好是utf-8

// 下面这个方法演示不同字符在go语言中所占字节数是不同的
// 求字符串的len，实际是计算字符串中每个字符字节长度的总和
// 如果想得到真正的字符串长度，通过以下两种方法
func GetStringLength() {
	var s string = "hello, 世界"
	fmt.Println("the string length is: ", len(s)) // 13
	// 通过 RuneCountInString 方法计算长度
	fmt.Println("RuneCountInString: ", utf8.RuneCountInString(s))
	// 通过将字符串转成数据类型rune的切片计算长度
	fmt.Println("Rune: ", len([]rune(s)))
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