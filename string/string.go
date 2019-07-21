package string

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

// golang中string底层是通过byte数组实现的。中文字符在unicode下占2个字节，
// 在utf-8编码下占3个字节， golang默认编码正好是utf-8

// 下面这个方法演示不同字符在go语言中所占字节数是不同的
// 求字符串的len，实际是计算字符串中每个字符字节长度的总和
// 如果想得到真正的字符串长度，通过以下两种方法
func GetStringLength() {
	s := "hello, 世界"
	fmt.Println("the string length is: ", len(s)) // 13
	// 通过 RuneCountInString 方法计算长度
	fmt.Println("RuneCountInString: ", utf8.RuneCountInString(s))
	// 通过将字符串转成数据类型rune的切片计算长度
	fmt.Println("Rune: ", len([]rune(s)))
}

// UTF8字符串作为交换格式是非常方便的，但是在程序内部采用rune序列可能更方便，
// 因为rune大小一 致，支持数组索引和方便切割
// 以下方法是rune和string的相互转换，在第一个Printf中的% x参数用于在每个十六进制数字前插入一个空格
func ToggleRuneString() {
	s := "プログラム"
	fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	r := []rune(s)
	fmt.Printf("%x\n", r) // "[30d7 30ed 30b0 30e9 30e0]"

	// 如果是将一个[]rune类型的Unicode字符slice或数组转为string，则对它们进行UTF8编码
	fmt.Println(string(r)) // "プログラム"

	// 将一个整数转型为字符串意思是生成以只包含对应Unicode码点字符的UTF8字符串
	fmt.Println(string(65))     // "A", not "65"
	fmt.Println(string(0x4eac)) // "京"

	// 如果对应码点的字符是无效的，则用'\uFFFD'无效字符作为替换：
	fmt.Println(string(1234567)) // "(?)"
}

// 测试一个字符串是否是另一个字符串的前缀
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// 字符串后缀测试
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// 包含子串测试
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

// 为一串字符串数字每隔三位添加一个逗号 如120030 => 120,030
// comma函数将在最后三个字符前位置将字符串切割为两个子串并插入逗号分隔符，
// 然后通过递归调用自身来计算出前面的子串
func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return Comma(s[:n-3]) + "," + s[n-3:]
}

// 非递归版本的comma方法，使用bytes.Buffer
// 注意：如果使用WriteByte方法要使用单引号，因为go语言中单引号表示rune类型，双引号表示字符串
func Comma2(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	p := n % 3
	for i := 0; i < n; i++ {
		if (i-p)%3 == 0 && i != 0 {
			buf.WriteString(",") // or buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}

// 完善comma函数，支持浮点数，使用递归处理
func Comma3(s string) string {
	var slice string
	index := strings.Index(s, ".")
	if index != -1 {
		slice = s[:index]
	} else {
		slice = s
	}
	// 将小数点前的整数提取出来slice进行操作
	n := len(slice)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	p := n % 3
	for i := 0; i < n; i++ {
		if (i-p)%3 == 0 && i != 0 {
			buf.WriteString(",") // or buf.WriteByte(',')
		}
		buf.WriteByte(slice[i])
	}
	// 拼接小数点和它后面的数字
	for i := index; i < len(s); i++ {
		buf.WriteByte(s[i])
	}
	return buf.String()
}
