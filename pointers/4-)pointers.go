package pointers

import "fmt"

// Pointers are just variables this is a proof not references
// value and ptr2 they are not same variable is just copy that's why address is different just copy of value
func destroy(value *int) {

	fmt.Printf("In Destroy value memory Adress : %p\t value points Address : %p\t value value : %d\n", &value, value, *value)
	value = nil
	fmt.Printf("In Destroy value=nil memory Adress : %p\t value points Address : %p\t value value : nil\n", &value, value)
}

func PointersAreNotReferences() {

	ptr2 := new(int)

	fmt.Printf("PTR2 memory Adress : %p\t PTR2 points Address : %v\t PTR2 value : %v\n", &ptr2, ptr2, *ptr2)

	*ptr2 = 10

	fmt.Printf("PTR2 memory Adress : %p\t PTR2 points Address : %p\t PTR2 value : %d\n", &ptr2, ptr2, *ptr2)

	destroy(ptr2)
	fmt.Println("after destroy function in main******************************************")
	fmt.Printf("PTR2 memory Adress : %p\t PTR2 points Address : %p\t PTR2 value : %d\n", &ptr2, ptr2, *ptr2)

}
