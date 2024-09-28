package model

import (
	"slices"
	"strconv"
	"strings"
)

type Student struct {
	Name        string
	Marks       map[string][]int
	AverageMark float64
	Classes     []string
}

func CreateStudent(name string, marks map[string][]int, classes []string) *Student {
	return &Student{Name: name, Marks: marks, Classes: classes}
}

func (s *Student) AddMark(class string, mark int) {
	s.Marks[class] = append(s.Marks[class], mark)
}

func (s *Student) AddClass(class string) {
	if slices.Contains(s.Classes, class) {
		return
	}
	s.Classes = append(s.Classes, class)
}

func (s *Student) CountAverageMark() {
	average := 0.0
	averageLessons := len(s.Classes)

	for _, lesson := range s.Classes {
		count := 0
		for _, mark := range s.Marks[lesson] {
			count += mark
		}

		average += float64(count) / float64(len(s.Marks[lesson]))
	}

	s.AverageMark = average / float64(averageLessons)
}

func (s *Student) MarksToString() string {
	builder := new(strings.Builder)

	for key, value := range s.Marks {
		builder.WriteString(key)
		builder.WriteString(": ")
		for index := range value {
			builder.WriteString(strconv.Itoa(value[index]))
			builder.WriteString(", ")
		}
		builder.WriteString("\n")
	}

	return builder.String()
}
