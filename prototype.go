package main

// importing this library allows you to output to console and collect user input
import "fmt"

func main() {
	fmt.Println("Hello World!")

	// if else statements
	// declare variable
	var number int
	fmt.Print("Enter a number: ")
	// Let user input number
	fmt.Scan(&number)

	fmt.Println(number)

	// check if number that user input is even or odd
	if number%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

	// use if else statements to determine grade using comparison operators based on user input
	var num float32
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	if num >= 90 {
		fmt.Println("Grade: A")
	} else if num >= 80 {
		fmt.Println("Grade: B")
	} else if num >= 70 {
		fmt.Println("Grade: C")
	} else if num >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}

	// use logical operator ||
	var temp int
	fmt.Print("Enter the temperature: ")
	fmt.Scan(&temp)

	if temp < 0 || temp > 100 {
		fmt.Println("Extreme temperature is detected.")
	} else {
		fmt.Println("Temperature is within normal range.")
	}

	// while loop
	x := 0
	for x < 5 {
		fmt.Println("Value of x is:", x)
		x++
	}

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Value of i is:", i)
	}

	// array
	var arr [5]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50

	fmt.Println("Array: ", arr)
	fmt.Println(arr[4])

	// slice
	slice := []int{10, 20, 30, 40, 50}

	fmt.Println("Slice:", slice)
	fmt.Println(slice[3])

	// Map
	m := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}

	fmt.Println("Map:", m)
}
