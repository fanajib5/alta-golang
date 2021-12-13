package main

import "fmt"

func main() {
	// LOOPING
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Printf("\n")

	// LOOPING WITH CONTINUE AND BREAK
	for i := 0; i < 5; i++ {
		if i == 1 {
			continue
		}
		if i > 3 {
			break
		}
		fmt.Println(i)
	}

	// LOOPING OVER STRING
	sentence := "Hello"
	for i := 0; i < len(sentence); i++ {
		fmt.Printf(string(sentence[i]) + "-")
	}
	fmt.Println("         >----------")
	for pos, char := range sentence {
		fmt.Printf("Character %c start at byte position %d\n", char, pos)
	}
}
