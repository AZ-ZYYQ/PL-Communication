package main

import "fmt"

func main() {
	//1. Verify the use of a chinese character declared as the variable name.
	var 字 int = 20
	fmt.Println(字)
	//Conclusion: Chinese character can be used as variable name

	//2. What it would be if the value list doesn't match the variable list?
	/*
		var A, B, C = "", 20
		fmt.Println(A)
		fmt.Println(B)
		fmt.Println(C)
	*/
	//The error occurs: Missing init expr for C

	//3. Can I specify different types for a set of variable in one declaration statement without using `= expression` ?
	/*
		var A float64, B int, C int
		fmt.Println(A)
		fmt.Println(B)
		fmt.Println(C)
	*/
	//The error occurs: unexpected comma at the end of statement
	//Conclusion: When declaring multiple variables without assigning an initialized value,
	//the data type must be specified clearly and the type of these variables must be the same.
	//For example: var a, b, c int

	//4. Verify the usage of short declaration in an outer block
	var arr [10]int

	for _, v := range arr {
		fmt.Println(v)
	}
	//in the loop, the variable `_` and `v` seem to be declared multiple times using := syntax, but it works well

	//5. Verify the multiple use of short declaration for declared variables.
	a, b := 10, 20
	a, c := 10, 30
	b, d := 40, 50
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	//The code above works well
	/*
		a, b := 10, 30
		fmt.Println(a)
		fmt.Println(b)
	*/
	//Error occurs: No new variable on left side of :=

}
