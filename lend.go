package main

import "fmt"
import "os"

func main() {
	args := os.Args[1:]

	if (len(args ) == 0) {
		interactiveMode()
	} else {
		parseArguments(args)
	}
}

func parseArguments(args []string) {
	switch {
		case len(args ) == 0 :
			fmt.Println("Supply at lease one argument")	
		case args[0] == "cat":
			category(args[1:])
		case args[0] == "human":
			human(args[1:])
		case args[0] == "item":
			item(args[1:])
		case args[0] == "switch":
			switchMode(args[1:])
		case args[0] == "lent":
			lent()
		// todo: make case for
		// - lend [item] to [human]
		// - lend [human] returned [item]
	}
}

func switchMode(args []string) {
	fmt.Println("switch between local and online")	
}

func interactiveMode() {
	fmt.Println("start interactive mode")	
}

func category(args []string) {
	fmt.Println("do something with a category")	
}

func human(args []string) {
	fmt.Println("do something with a human")	
}

func item(args []string) {
	fmt.Println("do something with an item")	
}

func lent() {
	fmt.Println("show list of items that are lent")	
}