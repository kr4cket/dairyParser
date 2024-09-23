package model

import (
	"bytes"
	"fmt"
	"strconv"
)

type Student struct {
	Name        string
	Marks       []int
	AverageMark float32
}

func CreateStudent(name string, marks []int, average float32) *Student {
	return &Student{Name: name, Marks: marks, AverageMark: average}
}

func (s *Student) AddMark(mark int) {
	s.Marks = append(s.Marks, mark)
}

func (s *Student) CountAverageMark() {
	var count = 0
	for _, mark := range s.Marks {
		count += mark
	}
	s.AverageMark = float32(count) / float32(len(s.Marks))
}

func (s *Student) GetInfo() {
	info := fmt.Sprintf("Name: %s\nMarks: %s\nAverage mark: %.2f \n", s.Name, IntegerArrayToString(s.Marks), s.AverageMark)
	println(info)
}

func IntegerArrayToString(array []int) string {
	var buffer bytes.Buffer
	for i := 0; i < len(array); i++ {
		buffer.WriteString(strconv.Itoa(array[i]))
		if i != len(array)-1 {
			buffer.WriteString(", ")
		}
	}

	return buffer.String()
}
