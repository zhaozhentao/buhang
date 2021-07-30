package types

import "strconv"

// StringToInt 将字符串转换为 int
func StringToInt(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
	}
	return i
}
