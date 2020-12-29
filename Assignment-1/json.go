package main

import (
	"encoding/json"
	"fmt"
)
type Human struct {
	Name    string
	Age     int
	Address string
}
func main() {

	human1 := Human{"xyz", 16, "nepal"}

	human_enc, err := json.Marshal(human1)

	if err != nil {

		fmt.Println(err)
	}

	fmt.Println(string(human_enc))

	human2 := []Human{
		{Name: "xyz", Age: 16, Address: "nepal"},
		{Name: "zyx", Age: 15, Address: "Pune"},
		{Name: "yxz", Age: 23, Address: "Bangalore"},
	}

	human2_enc, err := json.Marshal(human2)

	if err != nil {

		fmt.Println(err)
	}

	fmt.Println()
	fmt.Println(string(human2_enc))
}