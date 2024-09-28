package sorter

import (
	"diaryParser/pkg/model"
	"sort"
)

func Sort(studentsMap map[string]*model.Student) []*model.Student {
	var studentsSlice []*model.Student
	for key := range studentsMap {
		studentsSlice = append(studentsSlice, studentsMap[key])
	}

	sort.Slice(studentsSlice, func(i, j int) bool {
		if studentsSlice[i].AverageMark > studentsSlice[j].AverageMark {
			return true
		}

		if studentsSlice[i].AverageMark == studentsSlice[j].AverageMark {
			return studentsSlice[i].Name < studentsSlice[j].Name
		}

		return false
	})

	return studentsSlice
}
