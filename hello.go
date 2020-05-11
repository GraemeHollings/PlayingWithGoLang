package main

import (
	"fmt"
	"errors"
	"math"
)

func main() {
	/*GOLANG BASICS
	:= is the type inference operator, it infers that the value it is assigned is an int, so its cast as an int. 
	*/
	fmt.Println("------GOLANG BASICS------")
	fmt.Println("Hello world!")

	 x := 5
	 y := 10
	 z := 15
	
	sum := x + y + z

	if sum == 30 {
		fmt.Println("Our sum equals 30.")
	} else if sum != 30 {

		fmt.Println("Our sum does not equal 30.")
	} else
	{
		fmt.Println("nothing happened!")
	}
	
	//END OF GOLANG BASICS

	/*ARRAY STUFF
 	 Leave the [] empty so the array is not a fixed size array.
	 default declaration would be: var a [5] int
	 */
	 fmt.Println("------GOLANG ARRAYS------")
	 a := [] int {1, 2, 3, 4, 5}
	 a[2] = 7

	 //Append doesnt modify the original slice, it returns a new one. 
	 a = append(a, 13)

	 fmt.Println(a)

	//END OF ARRAY STUFF

	/*START OF MAPS
	Maps are like dictionaries or tables, a key and a value is given. 
	format for maps: map[type of keys] type of values.
	
	To make a map in go, use the built in make function, and give it a name.
	*/
	fmt.Println("------GOLANG MAPS------")
	vertices := make(map[string]int)
	vertices["triangle"] = 2;
	vertices["square"] = 3;
	vertices["circle"] = 4;

	fmt.Println(vertices)

	//use the same syntax to get a value from a particular key (lookup the value)
	fmt.Println(vertices["triangle"])

	//and use the delete function to delete something from the map
	delete(vertices, "square")

	fmt.Println(vertices)
	
	//END OF MAPS

	/* START OF LOOPS
	   The only loop in GoLang is the for loop. 
	*/
	fmt.Println("------GOLANG LOOPS------")
	for i := 0; i < 5; i++ {

		fmt.Println("The loop counter is:", i)

	}

	//The for loop also doubles as a while loop, so if you delete the variable declaration and the increment then you've got a while loop!
	fmt.Println("While loop example:")
	i := 0;
	for i < 5{
		fmt.Println("The while loop counter is:", i)
		i++
	}

	//iterate over each element over an array or slice
	arr := []string {"a", "b" , "c" }
	for index, value := range arr{
			fmt.Println("index:", index, "value:", value)
	}

	//the same thing can also be done with a map.
	m := make(map[string]string)
	m["a"] = "apple"
	m["b"] = "banana"

	for key, value := range m{
		fmt.Println("key:", key, "value:", value)
	}
	//END OF LOOPS

	//FUNCTIONS AND FUNCTION CALLING
	//Basic function calling example:
	result := calculateSum(2, 3)
	fmt.Println("Sum function:", result)

	sqrtResult, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	} else{
		fmt.Println("Sqrt function result:", sqrtResult)
	}

	//Important note, go does not use exceptions, thats why its so important that we return errors to assist us with debugging.

	//BEGINNING OF STRUCTURES:
	p := person{name: "Graeme", age: 20}
	fmt.Println(p)
	fmt.Println("Graemes age:", p.age)
	//END OF STRUCTURES.

	
}

//Return type in a function goes after the brackets
func calculateSum(x int, y int) int{
	return x + y
}

//Functions can also have multiple return types: 
func sqrt(x float64) (float64, error){
	if x < 0 {
		return 0, errors.New("undefined for negative numbers.")
	}

	return math.Sqrt(x), nil
}

//Structure example: 
type person struct{
	name string
	age int

}
