package main

import "fmt"

// myName := "shiv" // outside or globally not allowed

var myName = "Shiv"  // this is allowed

const getDetails string = "hey"  // public

func main() {

	// string variables 
    var username string = "Shiv"
	fmt.Println(username)
	fmt.Printf("variable is type of : %T\n", username)

	// boolean variable
	var isTrue bool = false
	fmt.Println(isTrue)
	fmt.Printf("variable is type of : %T\n", isTrue)

	// integer
	var smallVal uint8 = 233 // only can store upto 255
	fmt.Println(smallVal)
	fmt.Printf("variable is type of : %T\n", smallVal)

	// float
	var smallfloat float32 = 233.555555555545454848      // 233.55556
	fmt.Println(smallfloat)
	fmt.Printf("variable is type of : %T\n", smallfloat)

	// float
	var smallfloatwith64 float64 = 233.555555555545454848    // 233.55555555554545
	fmt.Println(smallfloatwith64)
	fmt.Printf("variable is type of : %T\n", smallfloatwith64)


	// default values nad some aliases
    var anothervar string   	// blank
	fmt.Println(anothervar)
	fmt.Printf("variable is type of : %T\n", anothervar)

	// implicit type

	var webSite = "hello.com"
	fmt.Println(webSite)

	// no var style

	numberOfUnits := 3999	
	fmt.Println(numberOfUnits)

	// using global variable

	fmt.Println(getDetails)
}