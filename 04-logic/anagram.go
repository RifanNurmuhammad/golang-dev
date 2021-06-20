package anagram

import (
	"sort"
)

type ListRune []rune

type GroupingStruct struct {
	original string
	group    int
}

func GetGroupAnagram(str []string) (grouping [][]string) {
	tempGroup := groupingAnagram(str)
	for _, v := range tempGroup {
		grouping = append(grouping, v)
	}
	return grouping
}

func groupingAnagram(str []string) map[string][]string {
	tempGroup := make(map[string][]string)
	for i := 0; i < len(str); i++ {
		value := sortSliceRune(str[i])

		if val, ok := tempGroup[value]; ok {
			val = append(val, str[i])
			tempGroup[value] = val
			continue
		}

		tempGroup[value] = []string{str[i]}
	}
	return tempGroup
}

func sortSliceRune(s string) string {
	var r ListRune = stringToRuneSlice(s)
	sort.Slice(r, func(k, v int) bool {
		return r[k] < r[v]
	})
	return string(r)
}

func stringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}
