package main

func isIsomorphic(s string, t string) bool {
	runeMap := make(map[rune]byte)
	reverseMap := make(map[byte]rune)
	for index, value := range s {
		iso, ok := runeMap[value]
		note, ok1 := reverseMap[t[index]]
		if !ok && !ok1 {
			runeMap[value] = t[index]
			reverseMap[t[index]] = value
			continue
		}

		if iso != t[index] || note != value {
			return false
		}
	}

	return true
}
