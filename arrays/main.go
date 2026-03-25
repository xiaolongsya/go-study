package main

import "fmt"

// type floatMap map[string]float64

// func (m floatMap) output() {
// 	fmt.Println(m)
// }

// func main() {
// 	userNames := make([]string, 2, 5)
// 	userNames[0] = "John"
// 	userNames[1] = "Jane"
// 	fmt.Println("User Names:", userNames)
// 	userNames = append(userNames, "Alice")
// 	userNames = append(userNames, "Bob")
// 	userNames = append(userNames, "Charlie")
// 	userNames = append(userNames, "Dave") // This will exceed the initial capacity and create a new underlying array
// 	fmt.Println("User Names:", userNames)

// courseRatings := make(floatMap, 5)
// courseRatings["Go Basics"] = 4.5
// courseRatings["Advanced Go"] = 4.8
// courseRatings.output()

// 	for index, value := range userNames {
// 		fmt.Printf("Index: %d, Value: %s\n", index, value)
// 	}

// }

type transformFunc func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	dNumbers := transformNumbers(&numbers, double)
	tNumbers := transformNumbers(&numbers, triple)
	fmt.Println("Original numbers:", numbers)
	fmt.Println("Doubled numbers:", dNumbers)
	fmt.Println("Tripled numbers:", tNumbers)

}

func transformNumbers(numbers *[]int, transform transformFunc) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
