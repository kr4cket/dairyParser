package app

import (
	"bufio"
	"diaryParser/pkg/model"
	"diaryParser/pkg/parser"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run(input string, output string) {
	var parseData []string

	if input != "" {
		parseData = ReadFile(input)
	} else {
		parseData = ReadConsole()
	}

	resultData := parser.Parse(parseData)

	if output != "" {
		SaveFile(output, resultData)
	} else {
		DisplayInConsole(resultData)
	}
}

func ReadFile(fileName string) []string {
	var result []string
	file, errs := os.Open(fileName)

	if errs != nil {
		fmt.Println("Failed to open file:", errs)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}

func ReadConsole() []string {
	var result []string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter student data in format: 'Name' 'class' 'mark'\nIf you want to stop write students, just press 'Enter'")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if text == "stop" {
			break
		}

		result = append(result, text)
	}

	return result
}

func SaveFile(fileName string, data []*model.Student) {
	file, errs := os.Create(fileName)
	if errs != nil {
		fmt.Println("Failed to create file:", errs)
		return
	}
	defer file.Close()

	builder := new(strings.Builder)

	for key := range data {
		builder.WriteString("Name: ")
		builder.WriteString(data[key].Name)
		builder.WriteString("\n")
		builder.WriteString("Marks:\n")
		builder.WriteString(data[key].MarksToString())
		builder.WriteString("Average mark: ")
		builder.WriteString(strconv.FormatFloat(data[key].AverageMark, 'f', 2, 64))
		builder.WriteString("\n")
		builder.WriteString("\n")
	}
	textToSave := builder.String()

	_, errs = file.WriteString(textToSave)
	if errs != nil {
		fmt.Println("Failed to write to file:", errs) //print the failed message
		return
	}
	fmt.Printf("Wrote to file %s.", fileName)
}

func DisplayInConsole(students []*model.Student) {
	for index := range students {
		fmt.Printf("Name: %s\nMarks:\n%sAverage mark: %.2f \n\n", students[index].Name, students[index].MarksToString(), students[index].AverageMark)
	}
}
