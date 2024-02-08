package main

import "fmt"

// func main() {
// 	// ARRAY EXAMPLE

// 	//eg1
// 	fmt.Println("Array samples in golang")
// 	// var arr1 [3]string
// 	var arr1 = [3]string{"Coffee", "Tea", "Other"}
// 	fmt.Println("arr1: ", arr1)

// 	//eg2 := means declaring and assigning at a time
// 	arr2 := [3]string{"Bus", "car", "bike"}
// 	fmt.Println("arr2: ", arr2)

// 	//eg3
// 	arr3 := arr2
// 	fmt.Println("arr3 : ", arr3)

// 	//eg4
// 	fmt.Println("length of the arr3: ", len(arr3))
// }

//SLICES

func main() {
	fmt.Println("Array samples in golang")
	//eg1
	var arr1 []string
	arr1 = []string{"car", "bus", "bike"}
	fmt.Println("arr1: ", arr1)
	// eg2
	var arr2 = []string{"cat", "dog", "Tiger"}
	fmt.Println("arr2: ", arr2)
	//eg3
	arr3 := []int{10, 20, 30}
	arr4 := []string{"mrudul", "achu"}
	fmt.Printf("value of arr3: %v , value of arr4: %v", arr3, arr4)
	fmt.Println("")

	// eg4
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("value of slice s = ", s)
	s = append(s, 6, 7, 8)
	fmt.Println("updated value of slice s = ", s)
}
