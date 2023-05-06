package main

import "fmt"

func main() {
	store := make(map[int]struct{})
	var command string
	var value int

	for {
		fmt.Print("(a/l/r/q) > ")
		fmt.Scanf("%s %d", &command, &value)

		// fmt.Println("You entered:", command)

		switch command {
		case "a":
			// If value is empty ask to enter value again
			if value == 0 {
				fmt.Println("Please enter a value")
				continue
			}

			if _, ok := store[value]; ok {
				fmt.Println("Already exists")
				continue
			}

			fmt.Println("Add", value)
			store[value] = struct{}{}
		case "l":
			fmt.Println("List")
			for k := range store {
				fmt.Println(k)
			}
		case "r":
			if len(store) == 0 {
				fmt.Println("No values to remove")
				continue
			}

			if _, ok := store[value]; !ok {
				fmt.Println("Does not exist")
				continue
			}

			fmt.Println("Remove", value)
			delete(store, value)
		case "q":
			fmt.Println("Bye")
			return
		}
	}
}
