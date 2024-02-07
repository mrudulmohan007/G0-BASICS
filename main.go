package main

import "fmt"

func main() {
	// ARRAY EXAMPLE

	//eg1
	fmt.Println("Array samples in golang")
	// var arr1 [3]string
	var arr1 = [3]string{"Coffee", "Tea", "Other"}
	fmt.Println("arr1: ", arr1)

	//eg2
	arr2 := [3]string{"Bus", "car", "bike"}
	fmt.Println("arr2: ", arr2)

	//eg3
	arr3 := arr2
	fmt.Println("arr3 : ", arr3)

	//eg4
	fmt.Println("length of the arr3: ", len(arr3))
}
