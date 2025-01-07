package leetcode

import "strconv"

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)

	for _, s := range strs {
		freq := make([]int, 26)
		for _, c := range s {
			freq[c-'a']++
		}

		key := ""
		for _, count := range freq {
			key += strconv.Itoa(count) + "#" //(Itoa) int to ascii - The number 5 becomes the character '5', which has an ASCII value of 53
		}

		anagrams[key] = append(anagrams[key], s)
	}

	result := make([][]string, 0, len(anagrams))
	for _, group := range anagrams {
		result = append(result, group)
	}

	return result
}
