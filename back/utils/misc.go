package utils

import (
    "context"
    "strings"
    "time"
)

func GetLast(tospl, sep string) string {
	return strings.Split(tospl, sep)[len(strings.Split(tospl, sep))-1]
}

func GetCtx() (context.Context, context.CancelFunc) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

    return ctx, cancel
}