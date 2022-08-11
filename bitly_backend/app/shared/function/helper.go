package function

import (
	"math/rand"

	"github.com/gomodul/envy"
)

func MakeUrl(url string) (result string) {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var n = envy.GetInt("random_string")
	if n <= 0 {
		n = 20
	}
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	result = string(b)
	return
}
