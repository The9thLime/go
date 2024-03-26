package main

import "fmt"

func main() {
	// Explicitly assigning data types to variables
	var (
		name string  = "Hello world!" // Explicit declaration
		num1 byte    = 254            // Byte is simply int8
		num2 int32   = -234729424
		num3 float32 = -424234.044
		TorF bool    = true
	)

	// Print variables with explicit data types
	fmt.Println(name, num1, num2, num3, TorF)

	// Implicitly declaring variables using the 'Walrus operator' or short variable declaration.
	// The compiler guesses the data type when variables are defined implicitly.
	numGuess := -10525.68

	fmt.Printf("%T\n", numGuess) // Returns float64
	// numGuess = "Hello" will return an error cause the compiler has already assigned the variable a suitable data type

	// Unassigned variables are assigned a default value

	var (
		str  string // Empty string ""
		num4 uint64 // 0
		bl   bool   // false
	)
	fmt.Println(str, num4, bl)

}
