package shortener

import (
	"math/rand"
	"net/url"
	"time"
)

// Using Bijective conversion between natural numbers (IDs) and short strings
// // ALPHABET the supported alphabet
// const ALPHABET = "0123456789abcdegfhijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"

// // BASE 64
// const BASE = 64

// func encode(num int64) string {
// 	str := ""
// 	for num > 0 {
// 		str = string(ALPHABET[num%BASE+1]) + str
// 		num = num / BASE
// 	}
// 	return str
// }

// func decode(str string) int64 {
// 	num := 0
// 	l := len([]rune(str))
// 	for i := 0; i < l; i++ {
// 		num = num*BASE + strings.Index(ALPHABET, string(str[i]))
// 	}
// 	return int64(num)
// }

func shortId() string {
	timestamp := int(time.Now().UnixNano() / int64(time.Millisecond))
	// timeString := strconv.Itoa(timestamp)
	str := NumberToBHex(timestamp, 36)

	str = str + RandomString(4)
	return str
}

// 10 進制轉換 n = ?進制
func NumberToBHex(num, n int) string {
	var num2char = "0123456789abcdefghijklmnopqrstuvwxyz"
	num_str := ""
	for num != 0 {
		yu := num % n
		num_str = string(num2char[yu]) + num_str
		num = num / n
	}
	return num_str
}

func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func isValidURL(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}
	return true
}
