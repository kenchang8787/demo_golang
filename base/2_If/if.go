/*
If
If else
If elseif
	這三個都跟認知上的用法差不多, 就不多贅述
*/
package main

import "fmt"

func main() {
	ifExample()
	ifElseExample()
	ifElseIfExample()
}

// If 範例
func ifExample() {
	fmt.Println("\n# If 範例")

	name := "Ken"

	if name == "Ken" {
		fmt.Println("Is Ken!")
	}
}

// If Else 範例
func ifElseExample() {
	fmt.Println("\n# If else 範例")

	name := "Ken"

	if name == "Ken Chang" {
		fmt.Println("Is Ken Chang!")
	} else {
		fmt.Println("Is not Ken Chang!")
	}
}

// If ElseIf 範例
func ifElseIfExample() {
	fmt.Println("\n# If else if 範例")

	name := "Ken Chang"

	if name == "Ken" {
		fmt.Println("Is Ken!")
	} else if name == "Ken Chang" {
		fmt.Println("Is Ken Chang!")
	} else {
		fmt.Println("Is not Ken or Ken Chang!")
	}
}
