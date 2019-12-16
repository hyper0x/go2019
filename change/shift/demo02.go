package main

import (
	"fmt"
)

func main() {
	operand1 := -2
	count1 := 1
	fmt.Printf("%d << %d: %d\n", operand1, count1, operand1<<count1)
	fmt.Printf("%d >> %d: %d\n\n", operand1, count1, operand1>>count1)

	operand2 := -1
	count2 := -8
	count2u := uint8(count2)
	fmt.Printf("%d << %d: %d\n\n", operand2, count2u, operand2<<count2u)

	shift3 := func(operand int, count int) (res int, err error) {
		defer func() {
			if p := recover(); p != nil {
				err = fmt.Errorf("%v", p)
			}
		}()
		res = operand << count
		return
	}
	operand3 := -1
	count3 := -8
	fmt.Printf("Operand: %d; Count: %d\n", operand3, count3)
	res, err := shift3(operand3, count3)
	fmt.Printf("Result: %d (ERROR: %v)\n", res, err)
}
