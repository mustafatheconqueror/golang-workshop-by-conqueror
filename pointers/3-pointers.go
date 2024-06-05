package pointers

import "fmt"

func DereferencingInPointers() {

	// create 1 int variable and 2 pointer and show in console and derefence it
	var (
		a int
		b *int
		c *int
	)

	a = 10
	b = &a
	c = b

	//Show The Console Values, and memory addresses.
	fmt.Printf("a value is :%v \t a memory address is:%p \n", a, &a)
	fmt.Printf("b value is :%v \t b memory address is:%p \n", b, &b)
	fmt.Printf("c value is :%v \t c memory address is:%p \n", c, &c)

	// Dereference the value of b and c pointer
	fmt.Printf("b valu is: %v", *b)

}
