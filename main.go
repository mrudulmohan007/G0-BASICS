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
// func main() {
// 	fmt.Println("Function example ->")
// 	result := Sum(10, 20)
// 	fmt.Println("Sum is : ", result)
// 	sum, sub := Calc(20, 10)
// 	fmt.Printf("value of sum is : %v\n value of sub is : %v\n", sum, sub)
// 	PrintName("Mrudul", "Mohan")

// }

// //sum of 2 numbers function
// func Sum(a, b int) int {
// 	return a + b
// }

// //multiple return values functions

// func Calc(a, b int) (int, int) {
// 	sum := a + b
// 	sub := a - b
// 	return sum, sub
// }

// //variadic functions ... -> ellipsis operator

// func PrintName(name ...string) {
// 	fmt.Println("My name is : ", name)
// }

//Loops
//1.infinite loop

// func main() {
// 	for {
// 		fmt.Println("Hello world")
// 	}

// }

//2.condition based loops

// func main() {
// 	i := 1
// 	for i < 5 {
// 		fmt.Println(i)
// 		i++
// 	}
// }

//3.counter based loops

// func main() {

// 	for i := 1; i < 5; i++ {
// 		fmt.Println(i)
// 	}

// }

//4.collection based loops

//eq.1- loops in array
// func main() {
// 	numbers := [4]int{100, 200, 300, 400}

// 	// for i, v := range numbers { //i=index, v = value
// 	// 	fmt.Println(i, v)
// 	// }

// 	for _, v := range numbers { // _ is unused variable
// 		fmt.Println(v)
// 	}

// }

//eq.2- loops in map

// func main() {
// 	letters := map[string]string{
// 		"a": "A",
// 		"b": "B",
// 		"c": "C",
// 	}

// 	for k, v := range letters {
// 		fmt.Println(k, v)
// 	}

// 	//to print only value,
// 	for _, v := range letters {
// 		fmt.Println(v)
// 	}

// }

//CONDITION STATEMENTS  / BRANCHING

// func main() {
// 	fmt.Println("Branching/Condition Statements")
// 	i := 5
// 	//1.If Statements
// 	if i < 6 {
// 		fmt.Println("i<6")
// 	} else if i > 6 {
// 		fmt.Println("i>6")
// 	} else {
// 		fmt.Println("Failed")
// 	}

// 	//2.switch statements
// 	//condition based switch
// 	switch i {
// 	case 3:
// 		fmt.Println("Value is 3")
// 	case 5:
// 		fmt.Println("Value is 5")
// 	default:
// 		fmt.Println("Default value")
// 	}
// 	//logical switch
// 	switch {
// 	case i < 6:
// 		fmt.Println("i<6")
// 	case i > 6:
// 		fmt.Println("i>6")
// 	default:
// 		fmt.Println("Default value")
// 	}

// 	//3.goto statement
// 	c := 8
// 	if c > 7 {
// 		fmt.Println("c is > 7")
// 		goto mylabel
// 	}

// mylabel:
// 	if c > 7 {
// 		fmt.Println("after label : c>7")
// 	}

// }

//OOP IN GO

// type Number int //custom data type

// func (num Number) is_even() bool {     -> //method receiver function
// 	return num%2 == 0
// }

// func main() {
// 	fmt.Println("Oops demo in go")
// 	var num Number
// 	num = 10
// 	fmt.Println(num)
// 	fmt.Println(num.is_even())
// }

// type Car struct {
// 	Name  string
// 	Model string
// 	Year  int
// }

// func (c Car) print() string {         //method receiver function
// 	return c.Name + "-" + c.Model

// }
// func main() {
// 	fmt.Println("Oops demo in go")
// 	minicooper := Car{
// 		Name:  "MC",
// 		Model: "Latest-2021",
// 		Year:  2021,
// 	}
// 	fmt.Println(minicooper)
// 	fmt.Println(minicooper.print())
// }

//INTERFACE(not complete , so dont refer)

// type PetrolEngine struct {
// 	Name string
// }

// func (e PetrolEngine) Start() {
// 	fmt.Println("Starting ......", e.Name)
// }

// type GasEngine struct {
// 	Name string
// }

// func (e GasEngine) Start() {
// 	fmt.Println("Starting ......", e.Name)
// }

// func main() {
// 	fmt.Println("interfaces example : ")
// 	waganorEngine := PetrolEngine{
// 		Name: "waganor",
// 	}
// 	waganorEngine.Start()

// 	gasEngine := GasEngine{
// 		Name: "gas",
// 	}
// 	gasEngine.Start()
// }

//interface example::

type Runner interface {
	Run()
}

type Dog struct {
	Name string
}

func (d Dog) Run() {
	fmt.Println(d.Name, "is running")
}

type Cat struct {
	Name string
}

func (c Cat) Run() {
	fmt.Println(c.Name, "is running")
}

func main() {
	var runner Runner
	runner = Dog{Name: "Australian shepherd"}
	runner.Run()
	runner = Cat{Name: "persian cat"}
	runner.Run()
}
