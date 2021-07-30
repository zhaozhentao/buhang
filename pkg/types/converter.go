package types

import "strconv"

// StringToInt 将字符串转换为 int
func StringToInt(str string) uint64 {
	i, err := strconv.ParseUint(str, 10, 64)

	if err != nil {
	}

	return i
}

func StringToFloat(str string) float64 {
	f, err := strconv.ParseFloat(str, 10)
	if err != nil {
	}
	return f
}
