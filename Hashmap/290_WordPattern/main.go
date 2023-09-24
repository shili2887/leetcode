package main

import "strings"

func wordPattern(pattern string, s string) bool {
	patternMap := make(map[rune]string)
	sMap := make(map[string]rune)
	sSlice := strings.Split(s, " ")
	if len(sSlice) != len(pattern) {
		return false
	}
	for index, value := range pattern {
		pp, ok := patternMap[value]
		ss, ok1 := sMap[sSlice[index]]
		if !ok && !ok1 {
			patternMap[value] = sSlice[index]
			sMap[sSlice[index]] = value
			continue
		}

		if pp != sSlice[index] || ss != value {
			return false
		}
	}

	return true
}
