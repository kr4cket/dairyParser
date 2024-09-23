package parser

import (
	"diaryParser/pkg/model"
	"sort"
	"strconv"
	"strings"
)

func StartSync(parseData []string) {
	students := make(map[string]*model.Student)
	var keys []string

	for _, value := range parseData {
		var data = strings.Split(value, " ")
		mark, _ := strconv.Atoi(data[1])
		_, value := students[data[0]]

		if !value {
			keys = append(keys, data[0])
			students[data[0]] = model.CreateStudent(data[0], []int{mark}, 0)
		} else {
			students[data[0]].AddMark(mark)
		}
	}
	sort.Strings(keys)

	for _, value := range keys {
		students[value].CountAverageMark()
		students[value].GetInfo()
	}
}
