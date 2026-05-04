package main

import "fmt"



func makeCoffe (kind string){
	fmt.Printf("Making %s coffee...... \n",kind)
}



func main() {
	// var name string = "asif" // normal way
	// name := "asif" //short way | most used in real application
	// name string: = "asif" //wrong way
	// var name = "asif"
    
	//group variable. | call multiple variable
	// var (
	// 	name string
	// 	age int
	// )

	//fmt.Println(name, age)
	// name = "asif"
	// age = 20

	// var x, y int; // short way
	// x= 25;
	// y=30;
	// var x, y int = 20, 30 // short way
	// const pi  = 3.1416
	// fmt.Println(x, y)
	// fmt.Println(pi);


	//data type in go
	// userName := "Asif"  // string
	// userAge := 25       // int
	// userHeight := 5.9   // float
	// userIsStudent := true // bool
	// fmt.Println(userName, userAge, userHeight, userIsStudent)

	//zero value of go
	// var age int;
	// fmt.Println(age); // number type zero value 0
	// var name string;
	// fmt.Println("this is string",name) // string type zero value ""
	// var isAdmin bool
    // fmt.Println("this is bool",isAdmin) // bool type zero value false
	// var score float64;
	// fmt.Println("this is float",score) // float type zero value 0
	// var score2 float32;
	// var myName string = "asif"
	// var age int = 25
	// var rating float64 = 4.5
	// // fmt.Printf("My name is %s and im %d years old and my rating is %.2f", myName, age, rating)
	// formattedString:= fmt.Sprintf("My name is %s and im %d years old and my rating is %.2f", myName, age, rating)
    
	// fmt.Println(formattedString)

	//function creationg
	makeCoffe("black")
	makeCoffe("Latte")
	

}

