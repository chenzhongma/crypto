package main

import (
	"math/big"
	"fmt"
)

func main() {

	var b1 big.Int
	var b2 big.Int
	var b3 big.Int
	var b4 big.Int

	b1.SetString("10000", 10)
	b2.SetString("20000", 10)

	fmt.Printf("b2 bytes : %v\n", b2.Bytes())

	b4.Add(&b1, &b2)
	fmt.Printf("b4 : %d\n", b4.Int64())

	s1 := []byte("4000")
	b3.SetBytes(s1)

	b3.Add(&b1, &b2)

	fmt.Printf("b3 : %v\n", b3)
	fmt.Printf("b3 : %v\n", b3.Int64())

}
