package app

import (
	"bufio"
	"diaryParser/pkg/parser"
	"fmt"
	"os"
)

func Run() {
	var parseData = ReadFile()
	println("===================================SYNC FUNCTION===================================\n")
	parser.StartSync(parseData)
	println("===================================ASYNC FUNCTION===================================\n")
	parser.StartAsync(parseData)
}

func ReadFile() []string {
	var result []string
	file, err := os.Open("src/students.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}
