package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome to Golang Todo App")
	http.HandleFunc("/", HelloUser)
	http.ListenAndServe(":8000", nil)
}

func HelloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user. Welcome to your todolist!"
	fmt.Fprintln(writer, greeting)
}

// The following variables are global variables
var taskOne = "Finish Go Crash Course"
var taskTwo = "Finish Algo trading Course"

//taskTwo := "Finish Algo trading course" This is another way of declaring variables

func main_initial() {
	var maritimeCode = []string{taskOne, taskTwo, "Charlie", "Delta", "Echo"} //This is a slice. Slice's can have any number of elements. Strings can be replaced with variables
	var goArray = [2]string{taskOne, taskTwo}                                 //This is an Array. Arrays in Go have fixed number of element
	fmt.Println(taskOne)
	fmt.Println(taskTwo)
	fmt.Println("Setup Algo Trades")
	fmt.Println(714) //Numbers

	fmt.Println(maritimeCode)
	fmt.Print(goArray)
	fmt.Println("")
	//For loop in GO
	fmt.Println("Loops")
	for _, code := range maritimeCode {
		fmt.Println(code)
	}
	for index, code := range maritimeCode {
		fmt.Println(index+1, ".", code)
		fmt.Printf("%d. %s\n", index+1, code)
	}
	printTasks(maritimeCode) //Calling the function

	maritimeCode = addTask(maritimeCode, "Foxtrot")
	printTasks(maritimeCode)
}

func printTasks(maritimeCode []string) {
	fmt.Println("Loops")
	for _, code := range maritimeCode {
		fmt.Println(code)
	}
	for index, code := range maritimeCode {
		fmt.Println(index+1, ".", code)
		fmt.Printf("%d. %s\n", index+1, code)
	}
}

func addTask(code []string, newCode string) []string {
	var updatedCodes = append(code, newCode)
	return updatedCodes
	// printTasks(updatedCodes)
}
