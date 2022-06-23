package main

import (
	"fmt"
	"example/greetings"
)


func main(){
	//get greetings message and print it.
	message := greetings.MyHello("Peggy")
	fmt.Println(message)
}