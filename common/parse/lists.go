package parse

import (
	"slices"
	"strconv"
	"strings"
)

// StringToIntList takes a string and separator and returns a list of ints
func StringToIntList(s string, sep string) ([]int, error) {
	var err error
	empty := func(s string) bool { return s == "" }
	strList := slices.DeleteFunc(strings.Split(s, sep), empty)
	intList := make([]int, len(strList))
	for i, str := range strList {
		intList[i], err = strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			return nil, err
		}
	}
	return intList, nil
}
