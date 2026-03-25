package main

import (
	"fmt"
)

type Products struct {
	name   string
	number int
}

func main() {
	hobbies := [3]string{"Cooking", "Traveling", "Gaming"}
	fmt.Println("Hobbies:", hobbies[0], hobbies[1], hobbies[2])
	fmt.Println("First hobby:", hobbies[0])
	newHobbies := hobbies[1:3]
	fmt.Println("Second and third hobbies:", newHobbies[0], newHobbies[1])
	// newHobbies2 := hobbies[0:2]
	newHobbies2 := hobbies[:2]
	fmt.Println("First and second hobbies:", newHobbies2[0], newHobbies2[1])
	newHobbies3 := newHobbies2[1:3]
	fmt.Println("Second and third hobbies:", newHobbies3[0], newHobbies3[1])

	newArray := []string{"Learn Go", "Build a project"}
	fmt.Println("Course goals:", newArray[0], newArray[1])
	newArray[1] = "Master Go"
	newArray = append(newArray, "Contribute to open source")

	products := []Products{
		{name: "Laptop", number: 1},
		{name: "Phone", number: 2}}
	fmt.Println("Products:", products[0].name, products[0].number, products[1].name, products[1].number)
	products = append(products, Products{name: "Tablet", number: 3})
	fmt.Println("Products after adding a new one:", products[0].name, products[0].number, products[1].name, products[1].number, products[2].name, products[2].number)

	prices := []float64{999.99, 499.99}
	discountedPrices := []float64{899.99, 399.99}
	prices = append(prices, discountedPrices...)
	fmt.Println("Prices after appending discounted prices:", prices)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
