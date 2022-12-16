package utils

import (
	"math/rand"
	"strings"
	"time"
)

var (
	lowerCharSet = "abcdedfghijklmnopqrst"
	upperCharSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberSet    = "0123456789"
	allCharSet   = lowerCharSet + upperCharSet + numberSet
	allLength    = len(allCharSet)
	length       = 8
)

func init() {
	rand.Seed(time.Now().Unix())
}

func GeneratePassword() string {
	var password strings.Builder

	for i := 0; i < 1; i++ {
		random := rand.Intn(len(lowerCharSet))
		password.WriteByte(lowerCharSet[random])
	}

	for i := 0; i < 1; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteByte(upperCharSet[random])
	}

	for i := 0; i < 1; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteByte(numberSet[random])
	}

	for i := 0; i < length-3; i++ {
		random := rand.Intn(allLength)
		password.WriteByte(allCharSet[random])
	}

	inRune := []rune(password.String())

	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})

	return string(inRune)
}
