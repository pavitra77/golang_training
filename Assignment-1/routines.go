package main

import "fmt"

func f(m int) 
{
	for i := 0; i < 5; i++ 
{
		fmt.Println(m, ":", i)
	}
}

func main() {
	go f(1)
	var input string
	fmt.Scanln(&input)
}