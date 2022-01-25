package main

import (
	"fmt"
	"math"
	"strings"
)

var myarr []float64

func encryptMul(str string, key int) {
	for _, char := range str {
		fmt.Println(int(char))
		myarr = append(myarr, math.Round(math.Mod(float64((int(char)-97)*key), 26)))
	}
	fmt.Println(myarr)
}
func decryptMul(myarr []float64, key int) {
	var decr []rune
	final_result := ""
	for _, i := range myarr {

		decr = append(decr, rune(math.Round(math.Mod(i*float64(findInverse(key)), 26))+97))
	}
	for _, i := range decr {
		final_result = final_result + string(i)

	}

	fmt.Println("result after decryption", final_result)
}

func findInverse(num int) int {
	for i := 0; i <= 26; i++ {
		if math.Mod(float64(num*i), 26) == 1 {
			return i
		}
	}
	return num
}
func main() {
	str := "neha"
	key := 3
	result := strings.ReplaceAll(str, " ", "")
	println(result)
	encryptMul(result, key)
	decryptMul(myarr, key)
}
