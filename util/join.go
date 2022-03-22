package util

import "bytes"

func Join(strs []string) string {
	// 声明字节缓冲
	stringBuilder := bytes.Buffer{}
	// 把字符串写入缓冲
	for _, v := range strs {
		stringBuilder.WriteString(v)
	}
	// 将缓冲以字符串形式输出

	return stringBuilder.String()
}

func JoinLike(str string) string {
	// 声明字节缓冲
	stringBuilder := bytes.Buffer{}
	// 把字符串写入缓冲
	stringBuilder.WriteString("%")
	stringBuilder.WriteString(str)
	stringBuilder.WriteString("%")
	// 将缓冲以字符串形式输出

	return stringBuilder.String()
}
