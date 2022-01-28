package main

import (
	"fmt"
	"math"
)

func main() {
	p := 5.0
	q := 18.0

	senderPrivateKey := 4.0

	x := math.Mod(math.Pow(p, senderPrivateKey), q)

	receiverPrivateKey := 8.0

	y := math.Mod(math.Pow(p, receiverPrivateKey), q)

	ksender := math.Mod(math.Pow(y, senderPrivateKey), q)

	kreceiver := math.Mod(math.Pow(x, receiverPrivateKey), q)
	fmt.Println(ksender)
	fmt.Println(kreceiver)

}
