package main

import (
	"fmt"
	"time"
)

// main function, started point of go
func main() {
  // print out directly
	fmt.Println("Hello, World -- directly")

  // call printMsg() to print out 
  printMsg("Hello, World -- call function")

  // define a variable to store a string
  var str = "Hello, Gavin & Dylan -- call function -- "

	//get current time and set in str2 variable
  var str2 = getCurrentTime()
  // pass the variable to printMsg(), print out string with timestamp
  printMsg(str + str2)

}

// print out a string 
func printMsg(msg string){
 fmt.Println(msg)
}

// get current time
func getCurrentTime() string{
	var currentTime = time.Now()
 	return currentTime.Format("2006-01-02 15:04:05")
}

