package random

import (
	"math/rand"
	"time"
)

const (
	UPPER_LETTER     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	UPPER_LETTER_LEN = len(UPPER_LETTER)
	LOWER_LETTER     = "abcdefghijklmnopqrstuvwxyz"
	LOWER_LETTER_LEN = len(LOWER_LETTER)
	NUMBER           = "0123456789"
	LETTER           = UPPER_LETTER + LOWER_LETTER
	LETTER_LEN       = len(LETTER)
	CHAR             = NUMBER + UPPER_LETTER + LOWER_LETTER
	CHAR_LEN         = len(CHAR)
)

func Int(min, max int) int {
	source := rand.NewSource(time.Now().UnixNano())
	return min + rand.New(source).Intn(max-min)
}

func Char() string {
	rint := Int(0, CHAR_LEN)
	return CHAR[rint : rint+1]
}

func String(length int) string {
	return strs(Char, length)
}

func Letter() string {
	rint := Int(0, LETTER_LEN)
	return LETTER[rint : rint+1]
}

func Letters(length int) string {
	return strs(Letter, length)
}

func strs(fn func() string, length int) string {
	if 0 >= length {
		panic("lenght must be greater than zero")
	}

	str := ""
	for i := 0; i < length; i++ {
		str = str + fn()
	}

	return str
}
