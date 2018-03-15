package kmformat

import (
	"strconv"
)

func KFormat(n uint64) string {
	if n <= 999 {
		return parseNormal(n)
	}
	return parseK(n)
}

func MFormat(n uint64) string {
	if n <= 999999 {
		return parseNormal(n)
	}
	return parseM(n)
}

func Format(n uint64) string {
	if n <= 999 {
		return parseNormal(n)
	}
	if n <= 999999 {
		return parseK(n)
	}
	return parseM(n)
}

func parseNormal(n uint64) string {
	return strconv.FormatUint(n, 10)
}

func parseK(n uint64) string {
	temp := float64(n) / 1000
	return strconv.FormatFloat(temp, 'f', 2, 64) + "k"
}

func parseM(n uint64) string {
	temp := float64(n) / 1000000
	return strconv.FormatFloat(temp, 'f', 2, 64) + "m"
}
