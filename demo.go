package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unsafe"
	"time"
	"sync"
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
	funcsInGo()
	structsInGo()
	interfacesInGo()
	errorHandlingGo()
	genericsInGo()
	pointersInGo()
	goRoutines()
	mutexInGo()
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

func funcsInGo() {
	// Similar to any other statically typed language, need to specify the data types for the func parameters and the return
	addFunc := func(x int, y int) int {
		return x + y
	}
	fmt.Println(addFunc(2, 3))

	// Functions in Go can return multiple data values, unlike C

	splitFullName := func(fullname string) (string, string, error) {
		names := strings.Fields(fullname)
		if len(names) < 2 {
			return "", "", errors.New("The given name cannot be split")
		}
		return names[0], names[1], nil
	}
	first, last, err := splitFullName("John Doe")
	if err == nil {
		fmt.Printf("First name: %s, Last name: %s\n", first, last)
	}

	// Passing functions as arguments to other functions

	squareFunc := func(num int) int { return num * num }

	operationHandler := func(callable func(int) int, value int) int {		
		return callable(value)
	}
	number := 5
	numberSquare := operationHandler(squareFunc, number)
	fmt.Printf("The square of %d is %d\n", number, numberSquare)

	// Currying in Go
	func1 := func(word string) func(string) string {
		return func(word2 string) string {
			return word + " " + word2
		}
	}
	chain := func1("Hello")
	w1 := "world"
	w2 := "John"
	fmt.Println(chain(w1))
	fmt.Println(chain(w2))

	// The spread operator in Go for passing in a variable number of arguments

	sumIter := func(numbers ...int) int {
		sum := 0
		for _, num := range numbers {
			sum += num
		}
		return sum
	}
	fmt.Printf("The sum of 4, 5, 6 is %d\n", sumIter(4, 5, 6))

	// How to spread an existing array into arguments
	numsForSum := []int{1,2,3,4,5}
	fmt.Printf("The sum of %d numbers is %d\n", numsForSum, sumIter(numsForSum...))

	// Named return variables

	mulIter := func(numbers ...int) (result int) {
		// result is automatically initialised as 0
		// If a different starting value is needed then only re-initialise it
		result = 1
		for _, num := range numbers {
			result *= num
		}
		return
	}
	// No need to write "return result" - This is implied based on the function definition
	fmt.Printf("Product of 3x4x5 is %d\n", mulIter(3, 4, 5))
}

// Example struct
type Person struct {
	name string
	age int8
	subjects []Subject
}

// Getters and setters
func (p Person) getName() string {
	return p.name
}

func (p *Person) setName(newName string) {
	p.name = newName
	fmt.Println("Name changed to " + newName)
}

func (p *Person) addSubject(newSubject Subject) {
	p.subjects = append(p.subjects, newSubject)
}

type Subject struct {
	subjectName string
	gpa float32
}

func structsInGo() {
	// Struct declaration syntax

	person := Person{"John", 17, []Subject{{"Maths", 3.5}}}
	// Can also initialise structs as key-value pairs name="John", age=17
	
	fmt.Printf("The person's name is %s\n", person.getName())

	person.setName("Jimmy")
	fmt.Printf("Person's name changed to %s\n", person.getName())

	// Composition: struct within a struct
	person.addSubject(Subject{"English", 3.0})
	fmt.Println(person.subjects)
}

type Speaker interface {
	Speak() string
}

// Create objects that will implement this interface
type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "Woof! My name is " + d.Name
}

type Robot struct {
    Model string
}

func (r Robot) Speak() string {
    return "Beep boop. Model " + r.Model + " active."
}

func interfacesInGo() {
	// No verbose syntax for implementing interfaces in Go, like Java
	// Interface implementation is implicit rather than explicit in Go.
	// As long as you're using all the methods defined in the interface

	introduceOneself := func(s Speaker) {
		fmt.Println(s.Speak())
	}

	dog := Dog{"Bruno"}
	robot := Robot{"C3PO"}

	introduceOneself(dog)
	introduceOneself(robot)
}

func deferredFunc() {
	fmt.Println("This is a deferred function executed before a panic.")
	// The error can be captured like a try except block using the recover() method
	r := recover() // The error message gets stored in this variable 
	if r == nil {
		fmt.Println("There is no error")
	} else {
		fmt.Println(r)
		// Code to handle gracious termination can go here
		fmt.Println("Handling error graciously for expected error. Unexpected errors allowed to propagate.")
	}
}

func errorHandlingGo() {
	// panic is the reserve word used to raise an exception and stop further code execution
	// If a panic occurs during some form of IO, then how would that be handled graciously in Go?
	// The answer is to use the 'defer' reserve word.
	// Defer ensures that certain functions are executed before the Go throws a panic
	// This is used like a context manager to graciously close files, or DB connections or sockets before the panic

	// defer deferredFunc()
	// panic("This caused a crash")
	// x := 5
	// y := 0
	// z := x / y
	// fmt.Println("This will only print if no panic " + fmt.Sprint(z))

	// Raising panics manually is not commonly done.
	// The recommended approach is to create an Error object from the errors module

	divideFunc := func(a int, b int) (float64, error) {
		if b == 0 {
			return 0.0, errors.New("Division by 0 error")
		}
		return float64(a) / float64(b), nil
	}

	result, err := divideFunc(5, 0)

	if err != nil {
		// Handle error
		fmt.Println("Following error occurred: " + fmt.Sprint(err))
	}
	// The line below will get executed, as error objects do not cause code execution to stop inherently:
	fmt.Println("No error occurred, code execution can proceed. " + fmt.Sprint(result))
}

func addGeneric[T int | float64 | uint](x T, y T) T {
	return x + y
}

// Rather than writing a long line of data type, can use an interface

type Numbers interface {
	int | float64 | uint
}

func subtractGeneric[T Numbers] (x T, y T) T {
	return x - y
}

// Using multiple letter representations for different generic data variables
func getValues[k comparable, v any](mp map[k]v) []v {
	values := []v{}

	for _, value := range mp {
		values = append(values, value)
	}
	return values
}

// generic structs

type GenericStruct[T any] struct {
	value T
}

func genericsInGo() {

	intSum := addGeneric(2, 3)
	floatSum := addGeneric(2.5, 3.5)
	uintSum := addGeneric(2, 4)

	fmt.Println(intSum, floatSum, uintSum)

	subgeneric := subtractGeneric(25, 15)
	fmt.Printf("25 - 15 = %d\n", subgeneric)

	// Example two
	exampleMap := map[string]int{
		"a": 100,
		"b": 200,
		"c": 300,
		"d": 400,
	}
	values := getValues(exampleMap)
	fmt.Println(values)

	struc1 := GenericStruct[string]{"hello"}
	struc2 := GenericStruct[int]{25}
	fmt.Println(struc1.value)
	fmt.Println(struc2.value)

}

type Book struct {
	id int
	name string
}

func (book *Book) setName(newName string) {
	book.name = newName
	// No need for '->' to dereference a struct attribute in Go 
}

func (book *Book) setId(newId int) {
	(*book).id = newId
	// Just as example to show that dereferencing a stuct in Go syntactically is the same as C
}

func pointersInGo() {
	// Pointers works very much like in C
	changeValue := func(num *int, newValue int) {
		*num = newValue
	}

	num1 := 5
	changeValue(&num1, 10) // & to access memory address of variable
	fmt.Printf("The value of num1 was changed to %d\n", num1)

	// Book example
	book := Book{1, "Title1"}
	book.setName("Title2")
	book.setId(25)
	fmt.Printf("Book details:\nName: %s\nID:%d\n", book.name, book.id)
}

func asyncAdd(x int, y int, ch chan int) {
	ch <- x + y
}

func definedSleep(count int, ch chan string) {
	time.Sleep(time.Duration(count) * time.Second)
	// fmt.Printf("Waking up after %d seconds\n", count)
	ch <- fmt.Sprintf("Sleep complete after %d seconds", count)
}

// ping is a SEND-ONLY channel (chan<-)
func ping(pings chan<- string, msg string) {
    pings <- msg
    // fmt.Println(<-pings) // <-- This would cause a COMPILE ERROR
}

// pong is a RECEIVE-ONLY channel (<-chan) for pings
// and a SEND-ONLY channel for pongs
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func goRoutines() {
	ch1 := make(chan int)
	go asyncAdd(10, 5, ch1)

	x := <-ch1
	fmt.Printf("The sum 10 + 5 calculated asynchronously: %d\n", x)

	// Multiple routines but only one call to the channel
	timeChannel := make(chan string)
	var result string
	// go definedSleep(4, timeChannel)
	// go definedSleep(3, timeChannel)
	// go definedSleep(2, timeChannel)
	// go definedSleep(1, timeChannel)

	// result = <-timeChannel
	// fmt.Println(result)
	// // Only the thread that completed first will be added to the result variable.

	// // Need to add multiple result lines to capture output from all the initiated go routines
	// result = <-timeChannel
	// fmt.Println(result) // 2 seconds

	// result = <-timeChannel
	// fmt.Println(result) // 3 seconds
	
	// result = <-timeChannel
	// fmt.Println(result) // 4 seconds
	
	// looping through the routines to capture the output for them all
	go definedSleep(4, timeChannel)
	go definedSleep(3, timeChannel)
	go definedSleep(2, timeChannel)
	go definedSleep(1, timeChannel)

	fmt.Println("Looping through to capture output of all routines:")
	for i:= 0; i < 4; i++ {
		result = <-timeChannel
		fmt.Println(result)
	}

	// Using select statement to capture outputs from different channels
	go asyncAdd(35, 5, ch1)
	go definedSleep(2, timeChannel)

	for j := 0; j < 2; j++ {
		select {
		case addResult := <-ch1:
			fmt.Println(addResult)
		case sleepResult := <-timeChannel:
			fmt.Println(sleepResult)
		}
	}
	// Without using the loop, only the output from the channel whose routine completed first would be captured.
	// select statement are asynchronous and do not block the execution of the main routine.

	// Directional channels
	
	// By default, all channels are bi-directional: data can be sent to the channel and received from the channel
	// Uni-directional channels can be defined, based on the positioning of the <- with the chan type
	// <-chan is a receive-only channel
	// chan<- is a send-only channel
	// Its considered good practice to call close(channelName) on a send-only channel.
	// Without closing the sender channel, the receiver channel will go into a deadlock, expecting more data to arrive.

	// Quick example below

	pings := make(chan string, 1)
    pongs := make(chan string, 1)

    ping(pings, "passed message")
    pong(pings, pongs)

    fmt.Println(<-pongs)

	// Buffered channels

	// By default, channels have a buffer capacity of 0.
	// This means that a receiving event cannot happen unless a send event preceded it SIMULTANEOUSLY.
	// A send event cannot happen, unless a receive event happens to clear out the previous send.

	// By passing in a buffer capacity, we allow for multiple non-blocking send and receive events to occur on the channel.
	// Send event will ONLY block, if the buffer is full.
	// Receive event will ONLY block if the buffer is nil.

	ch := make(chan string, 2)

    ch <- "first"
    ch <- "second"
    // The buffer is now full. 
    // Sending a third would cause a block/deadlock here.

    fmt.Println(<-ch)
    fmt.Println(<-ch)
}

type Counter struct {
	value int
	lock sync.Mutex
}

func incrementCounter(counter *Counter, wg *sync.WaitGroup) {
	defer wg.Done() // Notify the wait group that this function is complete
	
	counter.lock.Lock()
	defer counter.lock.Unlock() // Once the operation is succesfully complete or an error is thrown, the lock will always be released
	counter.value++
	fmt.Println(counter.value)
}

func mutexInGo() {
	// Multiple threads acting on a shared resource will create race conditions
	// Need to use a mutex
	counter := Counter{value: 0}

	// To manage multiple threads execution, need to create a task group
	// This is helpful when you want to await the execution of all tasks in the group
	// This is done using the WaitGroup method in the sync library

	var routineGroup sync.WaitGroup

	for i := 0; i < 20; i++ {
		routineGroup.Add(1) // Add takes a delta value
		go incrementCounter(&counter, &routineGroup)
	}
	routineGroup.Wait() // Waits for all 20 threads in the work group to be complete
	fmt.Println("Final value: ", counter.value)
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
