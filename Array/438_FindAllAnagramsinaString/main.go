package main

import "fmt"

func findAnagrams(s string, p string) []int {
	res := []int{}
	ls, lp := len(s), len(p)
	pMap := map[rune]int{}
	for _, pr := range p {
		pMap[pr]++
	}
	for i := 0; i+lp <= ls; i++ {
		charMap := map[rune]int{}
		words := s[i : i+lp]
		for _, word := range words {
			charMap[word]++
		}

		bMisMatch := false
		for _, pr := range p {
			if charMap[pr] != pMap[pr] {
				bMisMatch = true
			}
		}

		if !bMisMatch {
			res = append(res, i)
		}
	}

	return res
}

func findAnagramsDiff(s string, p string) []int {
	res := []int{}
	ls, lp := len(s), len(p)

	if lp > ls {
		return res
	}

	pMap := map[rune]int{}
	for i, pr := range p {
		pMap[rune(s[i])]++
		pMap[pr]--
	}

	differ := 0
	for _, c := range pMap {
		if c != 0 {
			differ++
		}
	}
	if differ == 0 {
		res = append(res, 0)
	}

	for i, char := range s[:ls-lp] {
		if pMap[char] == 1 {
			differ--
		} else if pMap[char] == 0 {
			differ++
		}

		pMap[char]--
		if pMap[rune(s[i+lp])] == -1 {
			differ--
		} else if pMap[rune(s[i+lp])] == 0 {
			differ++
		}

		pMap[rune(s[i+lp])]++

		if differ == 0 {
			res = append(res, i+1)
		}
	}

	return res
}

func main() {
	fmt.Println(findAnagramsDiff("cbaebabacd", "abc"))
	fmt.Println(findAnagramsDiff("abab", "ab"))
	fmt.Println(findAnagramsDiff("baa", "aa"))
}
