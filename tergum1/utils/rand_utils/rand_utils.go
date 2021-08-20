package rand_utils

import "math/rand"

const letters = "abcdefghijklmnopqrstuvwxyz"

func GetRandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
