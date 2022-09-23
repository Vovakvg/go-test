package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Println("Write your name: ")
	fmt.Scanf("%s\n", &name)
	fmt.Println("Write you age: ")
	fmt.Scanf("%d\n", &age)
	fmt.Printf("Your name is %s, your age is %d", name, age)
}
