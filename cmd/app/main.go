package main

import (
	"diaryParser/pkg/app"
	"fmt"
	"log"
	"os"
)

var commands = []string{"-o", "-output", "-f", "-file"}

func main() {
	input := ""
	output := ""

	args := os.Args[1:]

	for index := 0; index < len(args); index += 2 {
		if index+1 > len(args)-1 {
			log.Fatal(fmt.Sprintf("Command must have an argument"))
		}
		switch args[index] {
		case "-o", "-output":
			output = args[index+1]
		case "-f", "-file":
			input = args[index+1]
		default:
			log.Fatal("Unknown command")
		}
	}

	app.Run(input, output)
}
