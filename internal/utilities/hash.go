package utilities

import (
	"crypto/md5"
	"encoding/hex"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"

// HashFunc генерирует десятисимвольный URL
func HashFunc(text string) string {
	algorithm := md5.New()
	algorithm.Write([]byte(text))
	md5 := hex.EncodeToString(algorithm.Sum(nil))
	var str string
	for i, v := range md5[2:] {
		if i%3 == 0 {
			str += string(v)
		}
		i++
	}
	var nearestPrime rune = 67
	len := len(alphabet)
	var result string
	for _, v := range str {
		tmp := (v * nearestPrime) % rune(len)
		result = result + string(alphabet[tmp])
	}
	return result

}
