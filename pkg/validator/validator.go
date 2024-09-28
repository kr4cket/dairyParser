package validator

import (
	"fmt"
	"log/slog"
	"reflect"
	"slices"
	"strconv"
)

var lessonGoodness = []string{"physics", "math", "english"}

func IsLessonValid(lesson string, index int) bool {
	varType := reflect.TypeOf(lesson)

	if varType.Kind() != reflect.String {
		slog.Error(fmt.Sprintf("Line %d: Lesson must be type of string", index))
		return false
	}

	if !slices.Contains(lessonGoodness, lesson) {
		slog.Warn(fmt.Sprintf("Line %d: Lesson doesn't exist", index))
		return false
	}

	return true
}

func IsMarkValid(rawMark any, index int) bool {
	varType := reflect.TypeOf(rawMark)

	if varType.Kind() != reflect.String {
		slog.Error(fmt.Sprintf("Line %d: Couldn't convert this type to integer", index))
		return false
	}

	mark, err := strconv.Atoi(rawMark.(string))

	if err != nil {
		slog.Error(fmt.Sprintf("Line %d: Invalid variable to convert to integer", index))
		return false
	}

	if mark > 5 || mark < 1 {
		slog.Warn(fmt.Sprintf("Line %d: Mark must in [1..5]", index))
		return false
	}

	return true
}

func IsNameValid(rawName any, index int) bool {
	varType := reflect.TypeOf(rawName)

	if varType.Kind() != reflect.String {
		slog.Error(fmt.Sprintf("Line %d: Mark must be type of string", index))
		return false
	}

	return true
}

func IsValid(data []string, index int) bool {
	valid := true

	if !IsNameValid(data[0], index) {
		valid = false
	}

	if !IsLessonValid(data[1], index) {
		valid = false
	}

	if !IsMarkValid(data[2], index) {
		valid = false
	}

	return valid
}
