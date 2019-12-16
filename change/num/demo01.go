package main

import (
	"fmt"
)

func main() {
	numBin := 0b1110
	fmt.Printf("Binary integer: %d\n\n", numBin)

	numOct := 0o770
	fmt.Printf("Octal integer: %d\n\n", numOct)

	numHexInt := 0x10
	fmt.Printf("Hexadecimal integer: %d\n\n", numHexInt)

	numHexFP1 := 0x10p+1
	fmt.Printf("Hexadecimal floating point 1: %f\n", numHexFP1)
	numHexFP2 := 0x10p+2
	fmt.Printf("Hexadecimal floating point 2: %f\n", numHexFP2)
	numHexFP3 := 0x10p-1
	fmt.Printf("Hexadecimal floating point 3: %f\n\n", numHexFP3)

	numIm := 1.e+3i
	fmt.Printf("Imaginary: %v\n\n", numIm)

	NumSep1 := 3.14_15
	fmt.Printf("Number 1: %v\n", NumSep1)
	NumSep2 := 123_456_789
	fmt.Printf("Number 2: %v\n", NumSep2)
	NumSep3 := 1_2345_6789
	fmt.Printf("Number 3: %v\n", NumSep3)
	NumSep4 := 0x67_89
	fmt.Printf("Number 4: %v\n", NumSep4)
}
