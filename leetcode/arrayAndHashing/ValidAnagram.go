package leetcode

func isAnagram(s string, t string) bool {
    // if length of s and t are different the return false
    if len(s) != len(t){
        return false
    }

    charMap := make(map[rune]int)
        for _, c := range s {
        charMap[c]++
    }

    for _, c := range t{
        charMap[c]--

        if charMap[c] < 0{
            return false
        }
    }

    for _, count := range charMap {
        if count != 0 {
            return false
        }
    }

    return true
}