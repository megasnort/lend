package main

import "fmt"
import "os"

func main() {
	args := os.Args[1:]

	if (len(args ) == 0) {
		fmt.Println("Todo: activate mode where you do not have to prefix with program")
	} else {
		parseArguments(args)
	}
}

func parseArguments(args []string) {
	if (len(args ) == 0) {
		fmt.Println("Supply at lease one argument")		
	} else {
		fmt.Println(args)
	}
}