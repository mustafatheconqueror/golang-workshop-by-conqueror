package pointers

import "fmt"

// Never forget every variable has address and value

func PointersIntro() {

	//Create Pointer with var
	var (
		myPtr *int
	)

	//show in console
	fmt.Printf("myPtr values is:%v \t myPtr memory address: %p \n", myPtr, &myPtr)

	// Create Pointer with new() function
	myPtr2 := new(int)
	fmt.Printf("myPtr2 value is: %v \t myPtr2 memory address is : %p \n", myPtr2, &myPtr2)

	//show in console

}
