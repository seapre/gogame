# A simple calculator
* a calculator to add two numbers, or subtract, or multiply, or divide
* create a folder named as calc
* create main.go in the folder
* -- main.go --

```go
package main

import (
	"fmt"	
)

func main() {
	fmt.Println("Go calculator")
	cmd := readLine("Enter command: [a]dd, [s]ubtract, [m]ultiply, [d]ivide: ")
	fmt.Println(cmd)
}
```
* [how to create a simple calculator](https://www.youtube.com/watch?v=-GoCCBWGHhs)

* add more functions: add, subtract, multiply, divide
* scan user input two numbers 
* final main.go

```go
package main

import (
	"fmt"	
	"strconv"
)


func main() {
	fmt.Println("Go calculator")
	cmd := readLine("Enter command: [a]dd, [s]ubtract, [m]ultiply, [d]ivide: ")
	fmt.Println(cmd)
	if cmd == "a" {
		num1, num2 := getUserNumbers()
		result := num1 + num2
		sResult := fmt.Sprintf("%d", result )
		fmt.Println(sResult)
		
		} else if cmd == "s" {
		num1, num2 := getUserNumbers()
		result := num1 - num2
		sResult := fmt.Sprintf("%d", result )
		fmt.Println(sResult)
		} else if cmd == "m" {
		num1, num2 := getUserNumbers()
		result := num1 * num2
		sResult := fmt.Sprintf("%d", result )
		fmt.Println(sResult)
		} else if cmd == "d" {
		num1, num2 := getUserNumbers()
		result := float32(num1) / float32(num2)
		sResult := fmt.Sprintf("%f", result )
		fmt.Println(sResult)
	} else {
		fmt.Println("Invalid command")
	}
}
func readLine(message string) string{
	fmt.Print(message)
  var input string
	fmt.Scanln(&input)
	return input
}

func getUserNumbers() (int, int) {
	num1String := readLine("Enter first number: ")
  num1 ,err := strconv.Atoi(num1String)
	if err != nil {
		fmt.Println(err)
	}
	num2String := readLine("Enter second number: ")
  num2 ,err := strconv.Atoi(num2String)
	if err != nil {
		fmt.Println(err)
	}

	return num1, num2

}

```
* run it

```cmd
 > go build
 > ./calc.exe
```