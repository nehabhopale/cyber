package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/andrewarchi/gocipher/mod"
)

var myarr []float64
var key []int

func encryptAffine(str string) {
	for _, char := range str {

		myarr = append(myarr, math.Round(math.Mod(float64((int(char)-97)*key[0]+key[1]), 26)))
	}
	fmt.Println("array after encryption", myarr)

}
func decryptAffine(myarr []float64) {
	var decr []rune
	final_res := ""
	for _, i := range myarr {
		keyInv, _ := mod.Inverse(key[0], 26)

		decr = append(decr, rune(math.Round(math.Mod((i-float64(key[1]))*float64(keyInv), 26))+97))
	}
	for _, i := range decr {
		final_res += string(i)
	}

	fmt.Println("final_res after decryption", final_res)
}

func main() {
	key = append(key, 3)
	key = append(key, 4)
	str := "Neha"
	result := strings.ReplaceAll(str, " ", "")
	println("actual string", result)
	encryptAffine(result)
	decryptAffine(myarr)
}
