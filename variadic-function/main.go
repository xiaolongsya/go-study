package main

func main() {
	result := sumup(1, 2, 3, 4, 5)
	println("The sum is:", result)
	numbers := []int{6, 7, 8, 9, 10}
	anotherSum := sumup(numbers...)
	println("The sum of the slice is:", anotherSum)

}

func sumup(numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}
