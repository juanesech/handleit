package utils

import (
	"strings"
)

func GetLast(tospl, sep string) string {
	return strings.Split(tospl, sep)[len(strings.Split(tospl, sep))-1]
}
