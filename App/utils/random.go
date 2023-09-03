package utils

import "math/rand"

var runes = []rune("0987654321abcdefghiklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomUrl(size int) string {
	str := make([]rune, size)
	for i := range str {
		str[i] = runes[rand.Intn(len(runes))]
	}
	return string(str)
}
