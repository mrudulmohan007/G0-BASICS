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

// func main() {
// 	fmt.Println("Array samples in golang")
// 	//eg1
// 	var arr1 []string
// 	arr1 = []string{"car", "bus", "bike"}
// 	fmt.Println("arr1: ", arr1)
// 	// eg2
// 	var arr2 = []string{"cat", "dog", "Tiger"}
// 	fmt.Println("arr2: ", arr2)
// 	//eg3
// 	arr3 := []int{10, 20, 30}
// 	arr4 := []string{"mrudul", "achu"}
// 	fmt.Printf("value of arr3: %v , value of arr4: %v", arr3, arr4)
// 	fmt.Println("")

// 	// eg4
// 	s := []int{1, 2, 3, 4, 5}
// 	fmt.Println("value of slice s = ", s)
// 	s = append(s, 6, 7, 8)
// 	fmt.Println("updated value of slice s = ", s)
// }

//Maps in go
// func main() {
// 	fmt.Println("Golang with maps")
// 	temp := map[string]string{
// 		"key1": "value1",
// 		"key2": "value2",
// 	}
// 	fmt.Println(temp)

// 	m := map[string][]string{
// 		"coffee": {"Coffee", "Coffee-Late"},
// 		"tea":    {"chai", "Latte", "Chai-Latte"},
// 	}
// 	fmt.Println(m)
// 	fmt.Println(m["coffee"])

// 	//append
// 	m["others"] = []string{"milk", "water"}
// 	fmt.Println(m)

// 	//deleting
// 	delete(m, "tea")
// 	fmt.Println(m)

// }

//struct - on;y heterogenous datatype in go
//struct can be created in 2 ways

//1st method

// var Car struct {
// 	id   int
// 	name string
// }

// func main() {
// 	fmt.Println(" first Example of struct")
// 	//eg1
// 	Car.id = 1
// 	Car.name = "swift"
// 	fmt.Println(Car)
// 	fmt.Println(Car.name)

// }

//2nd method

// type Car struct {
// 	id      int
// 	name    string
// 	mileage float32
// }

// func main() {
// 	fmt.Println("2nd example of struct")
// 	car1 := Car{
// 		id:      1,
// 		name:    "Minicooper",
// 		mileage: 10,
// 	}
// 	fmt.Println("details of the first car :", car1)
// 	car2 := Car{
// 		id:      1,
// 		name:    "Swift",
// 		mileage: 20,
// 	}
// 	fmt.Println("details of the second car :", car2)

// }

//eg of a complicated struct

// func main() {
// 	fmt.Println("eg of complicated struct : ")
// 	type menuItem struct {
// 		name   string
// 		prices map[string]float64
// 	}

// 	//slice of menuItem struct
// 	menu := []menuItem{
// 		{
// 			name: "Coffee",
// 			prices: map[string]float64{
// 				"regular": 12.5,
// 				"large":   20,
// 			},
// 		},
// 		{
// 			name: "tea",
// 			prices: map[string]float64{
// 				"single": 10,
// 				"double": 15,
// 				"triple": 20,
// 			},
// 		},
// 	}
// 	fmt.Println(menu)

// }

//FUNCTIONS
func main() {
	fmt.Println("Function example ->")
	result := Sum(10, 20)
	fmt.Println("Sum is : ", result)
	sum, sub := Calc(20, 10)
	fmt.Printf("value of sum is : %v\n value of sub is : %v\n", sum, sub)
	PrintName("Mrudul", "Mohan")

}

//sum of 2 numbers function
func Sum(a, b int) int {
	return a + b
}

//multiple return values functions

func Calc(a, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}

//variadic functions ... -> ellipsis operator

func PrintName(name ...string) {
	fmt.Println("My name is : ", name)
}
