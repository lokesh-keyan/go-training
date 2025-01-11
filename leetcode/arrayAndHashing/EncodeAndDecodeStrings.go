package leetcode

import "strings"

type Codec struct {
}

func (codec *Codec) Encode(strs []string) string {
	res := strings.Builder{}
	for _, str := range strs {
		res.WriteByte(byte(len(str)))
		res.WriteString(str)
	}

	return res.String()
}

func (codec *Codec) Decode(strs string) []string{
	var count int
	var data string
	var ans []string
	for i := 0; i < len(strs); i += count + 1 {
		count = int(strs[i])
		data = strs[i+1:i+1+count]
		ans = append(ans, data)
	}

	return ans
}