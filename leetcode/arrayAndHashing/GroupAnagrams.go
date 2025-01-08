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

	result := make([][]string, 0, len(anagrams)) // here we are creating a slice of slices of strings with length 0 and capacity equal to the length of the map anagrams
	// make([]int, 3, 5) creates a slice with a length of 3 with 0 in three indices and a capacity of 5. The length is the number of elements referred to by the slice, 
	//	the capacity is the number of elements in the underlying array, counting from the first element in the slice.
	for _, group := range anagrams {
		result = append(result, group)
	}

	return result
}
