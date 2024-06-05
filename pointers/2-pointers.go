package pointers

import "fmt"

func PointersAreJustVariables() {

	//Crate 1 int variable and 1 pointer
	var (
		a int
		b *int
	)

	a = 10
	b = &a

	//Don't forget every variable has value and memory address. Show in console a and b
	fmt.Printf("a value is :%v \t a memory address is :%p \n", a, &a)
	fmt.Printf("bPtr value is: %v \t  bPtr memory address is :%p \n", b, &b)

}
