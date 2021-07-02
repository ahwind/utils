package system

import (
	"strconv"
)

func S2I(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 60)
	return i
}

func S2F(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
