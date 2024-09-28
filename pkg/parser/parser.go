package parser

import (
	"diaryParser/pkg/model"
	"diaryParser/pkg/sorter"
	"diaryParser/pkg/validator"
	"strconv"
	"strings"
)

func Parse(parseData []string) []*model.Student {
	students := make(map[string]*model.Student)
	var keys []string

	for index, value := range parseData {

		var data = strings.Split(value, " ")

		if !validator.IsValid(data, index+1) {
			continue
		}

		mark, _ := strconv.Atoi(data[2])

		class := data[1]
		name := data[0]
		_, value := students[name]

		if !value {
			keys = append(keys, name)
			students[name] = model.CreateStudent(data[0], map[string][]int{class: {mark}}, []string{class})
		} else {
			students[name].AddMark(class, mark)
			students[name].AddClass(class)
		}
	}

	for key := range students {
		students[key].CountAverageMark()
	}

	return sorter.Sort(students)
}
