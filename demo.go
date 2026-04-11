package main
import (
	"fmt"
	"unsafe"
	"math"
	"strconv"
)

func main() {
	fmt.Println("Hello world")
	variablesInGo()
	fmtBasics()
	arithmeticOperations()
	stringTypeConversion()
	goConditionals()
	loopsInGo()
	arraysInGo()
	slicesInGo()
	mapsInGo()
}

func variablesInGo() {

	// Explicit variable declaration:
	var num1 uint = 233
	fmt.Println(unsafe.Sizeof(num1))

	var num2 int = -233
	fmt.Println(unsafe.Sizeof(num2))
	// default size of int on my machine and Fedora 43 is 8 bytes

	var num3 float64 = 22.0 / 7.0
	fmt.Println(unsafe.Sizeof(num3))
	fmt.Println(num3)
	// Can use C syntax to format the string output:
	fmt.Printf("22 / 7 to two significant figures: %0.2f \n", num3)

	var char1 rune = 'h'
	fmt.Printf("%c\n", char1)
	// A rune is like the char datatype in C
	// In C, a char is 1-byte representing an ASCII character
	// In Go, a rune is 4-bytes representing a Unicode character

	// Implicit variable declaration:
	imp_int := 25
	imp_float := 34.7
	imp_bool := false
	imp_string := "hello"
	imp_binary := 0b110010

	fmt.Printf("%T \n", imp_int)
	fmt.Printf("%T \n", imp_float)
	fmt.Printf("%T \n", imp_bool)
	fmt.Printf("%T \n", imp_string)
	fmt.Printf("%T \n", imp_binary)

	// Implicit declarations can be type casted to desired data types:
	imp_int64 := int64(999)
	fmt.Printf("%T \n", imp_int64)
}

func fmtBasics() {
	// Println is similar to Python's implementation of print
	// \n is added by default to the end of the string

	// Printf is similar to C's implementation.
	// Need to use specific string formatters and add \n manually.
	num1 := 5
	num2 := 5.5
	bool1 := false
	string1 := "hello"
	char1 := 'g'

	fmt.Printf("%d\n", num1)
	fmt.Printf("%f\n", num2)
	fmt.Printf("%t\n", bool1)
	fmt.Printf("%s\n", string1)
	fmt.Printf("%c\n", char1)

	// Sprintf allows for a formatted string to be stored in a variable for future use
	pi := 22.0 / 7.0
	y := fmt.Sprintf("%0.2f", pi)
	fmt.Printf("%s\n", y)
}

func arithmeticOperations() {
	// Can only perform operations on identical data types
	// Need to typecast for differing data types

	// Nuance with int to string typecasting to watch out for:
	x := "hello"
	y := 2
	// incorrect := x + string(y) // This will concatenate the ASCII/Unicode representation of the integer 2
	// fmt.Println(incorrect)
	correct := x + fmt.Sprint(y) // This will concatenate the string literal of 2, which is the expected behaviour
	fmt.Println(correct)

	// Common functions in the math module:
	fmt.Println(math.Max(6, 80))
	fmt.Println(math.Min(6, 80))
	fmt.Println(math.Pow(4, 3))
	fmt.Println(math.Sqrt(9))
	fmt.Println(math.Round(4.25)) // rounds to 4
	fmt.Println(math.Round(4.75)) // rounds to 5
	fmt.Println(math.Ceil(4.25))
	fmt.Println(math.Floor(4.75))
}

func stringTypeConversion() {
	// Converting a string to other datatypes requires the strconv module
	// It contains a collection of conversion functions for different use cases

	stringNum := "1234"
	intNum, convError := strconv.Atoi(stringNum)
	fmt.Printf("String to int conversion using Atoi: %d\nErrors: %v", intNum, convError)
	
	// can also use ParseInt:
	intNum2, convError := strconv.ParseInt(stringNum, 10, 0)
	println(intNum2)
	
	// Converting a binary string to int requires using base 2:
	stringBin := "11010101"
	intBin, convError := strconv.ParseInt(stringBin, 2, 0)
	fmt.Printf("%s binary in integer form is %d\n", stringBin, intBin)

}

func goConditionals() {
	x := 9

	if x < 3 {
		fmt.Println("Less than three")
	} else if x < 8 {
		fmt.Println("Less than eight")
	} else {
		fmt.Println("Greater than or equal to 9")
	}

	state := "On"
	switch state {
	case "Off":
		fmt.Println("Machine is off,")
	case "On":
		fmt.Println("Machine is on")
	default:
		fmt.Println("Incorrect machine state")
	}
	// Switch cases do not require a break clause between cases. The break is implied.
	// If you wanted to enter another case after successfully meeting a case conditional
	// then explicitly write the "fallthrough" keyword after the case definition

	// Naked switch statement example:
	// Allows for conditional statement definition for each case
	switch {
	case x < 3:
		fmt.Println("Less than three")
	case x < 8:
		fmt.Println("Less than eight")
	default:
		fmt.Println("Incorrect machine state")
	}

	// Multiple case checks within a single case

	letter := 'h'
	switch letter {
	case 'a', 'b', 'c':
		fmt.Println("The letter is either a, b or c")
	default:
		fmt.Println("Its another letter in the English lexicon")
	}
	
	switch {
	case letter == 'a', letter == 'b', letter == 'c':
		fmt.Println("The letter is either a, b or c")
	default:
		fmt.Println("Its another letter in the English lexicon")
	}
}

func loopsInGo() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// The for loop definition is written line in C or JS, but without the paranthesis
	// Go doesn't have 'while' as a reserve word
	// While loops are created using 'for'
	j := 0
	for j < 5 {
		fmt.Printf("while loop iteration: %d\n", j)
		j++
	}

	// Looping through a string in Go:
	phrase := "Hello world"

	for letter := 0; letter < len(phrase); letter++ {
		fmt.Printf("Letter: %c and its type: %T\n", phrase[letter], phrase[letter])
	}
	// The above will not work for UTF-8 characters, only for ASCII, as the type is uint8
	for _, v := range phrase {
		fmt.Printf("Letter: %c and its type: %T\n", v, v)
	}
	// using range to iterate over a string casts each character as int32, ie, a rune.
}

func arraysInGo() {
	// Explicit declaration methods:
	// Declare size and name, initialise values later:
	var array1 [2]int32
	array1[0] = 2
	array1[1] = 8
	fmt.Println(array1)

	// Declare and initialise size and values:
	var array2 = [3]int32{1,2,3}
	fmt.Println(array2)

	// Implicit declaration of size based on the number of initialised elements in the array:
	// In C, its simply [], but in Go, need to use ellipsis [...]
	var array3 = [...]int32{1,2,3,4,5,6}
	fmt.Println("The length of array3 is " + fmt.Sprint(len(array3)))

	// Initialising specific indices within the array, others default to 0:
	var array4 = [5]int32{0:100, 4:200}
	fmt.Println(array4)

	// 2D arrays and an example of initialisation with walrus operator:
	array5 := [...][2]int32{{1, 2}, {2, 2}, {3, 2}}
	
	// Changing a row in the 2D array requires the new row to be initialised as an array:
	new_entry := [2]int32{10, 11}
	array5[0] = new_entry
	fmt.Println(array5)
}

func slicesInGo() {
	// Retrieve a slice of an array which can be stored as a separate variable
	// The slice is a pointer to the original array

	arr := [5]int{1,2,3,4,5}
	s1 := arr[1:3]
	fmt.Println(s1)

	s1[0] = 50
	fmt.Println(s1, arr) // Both arrays modified as its the same arr in memory

	// Concept of slice and capacity:
	fmt.Printf("The slice and its capacity: %d %d\n", s1, cap(s1))
	// The slice was started at index 1 of arr
	// The capacity of s1 is len(arr) - starting index of slice, 5 - 1 = 4
	// A slice of a static array CANNOT expand beyond the length of the static array

	// However, a slice instantiated from the start behaves like a dynamic array in Go:
	var slice = []int{0, 1}

	for i := 2; i < 15; i++ {
		slice = append(slice, i)
		fmt.Printf("Dynamic array: %d; Its size: %d\n", slice, cap(slice))
	}
	// Once the capacity is maxed out, the dynamic array's capacity is doubled, if malloc is successful.

	// Use the make method to create a slice array of a particular size to start

	customSlice := make([]int, 10) // Creates slice of size 10
	fmt.Println(customSlice)
}

func mapsInGo() {
	// Literal initialisation:
	map1 := map[string]int{"a": 1, "b": 2}
	fmt.Println(map1["a"])

	// Empty literal initialisations:
	map2 := map[string]int{}
	map2["one"] = 1
	map2["two"] = 2
	fmt.Println(map2)

	// Method 2:
	map3 := make(map[string]int) // Can pass optional parameter to declare capacity of map
	map3["one"] = 1
	map3["two"] = 2
	fmt.Println(map3)

	map3["three"] = 3

	// deleting a key from the map
	delete(map3, "two")
	fmt.Println(map3)

	// Check if a key exists in the map
	value, ok := map3["two"]
	fmt.Println("The two values returned when accessing a key from a map:")
	fmt.Println(value, ok)
	if !ok {
		fmt.Println("The key doesn't exist in the map.")
	}

	// Quick working example:
	divisibility := map[uint]uint{}
	maxNum := uint(100)

	for dividend := uint(1); dividend <= maxNum; dividend++ {
		for divisor := uint(1); divisor <= 5; divisor++ {
			if dividend % divisor == 0 {
				divisibility[divisor]++
				// All key values are initialised to zero, even if they don't exist.
				// So, no need to check if the key exists in this case before incrementing the value
			}
		} 
	}
	fmt.Println(divisibility)
}

// Video source of the tutorial: https://youtu.be/V-lI7AmusGs?si=ES7jOhBohipOllcB
// All go code needs to be part of a package.
// This is why we need package <name> at the top of every go file.
// Similar to C, every package needs a main() function.
// This is the entry point for the machine into the application.
// There are two way to run Go code:
// Compile the code to create a binary file, similar to C using the command: go build -o <bin-file-name> <go-file>
// Or, you can use the command: go run <go-file>
// This handles the build, run, teardown process into a single command.
