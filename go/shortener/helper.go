package shortener

import (
	"strconv"
	"strings"
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
	timeString := strconv.Itoa(timestamp)
	str := numberToBHex(timestamp, 36)

	str = str + string([]rune(timeString)[len(timeString)-4:len(timeString)])
	return str
}

func numberToBHex(num, n int) string {
	num2char := "0123456789abcdefghijklmnopqrstuvwxyz"
	num_str := ""
	for num != 0 {
		yu := num % n
		num_str = string(num2char[yu]) + num_str
		num = num / n
	}
	return strings.ToUpper(num_str)
}
