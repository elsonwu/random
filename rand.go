package random

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
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

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Int(min, max int) int {
	source := rand.NewSource(time.Now().UnixNano())
	return min + rand.New(source).Intn(max-min)
}

func Char() string {
	rint := Int(0, CHAR_LEN)
	return CHAR[rint : rint+1]
}

func String(length int) string {
	s := ""
	if length > 32 {
		var ts int = length / 32
		for i := 0; i < ts; i++ {
			s += String(32)
		}

		if tr := length % 32; tr > 0 {
			s += String(tr)
		}

		return s
	}

	s = Md5(strconv.Itoa(time.Now().Nanosecond()))
	if length == 32 {
		return s
	}

	return s[0:length]

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
