package tools

import (
	"strings"
)

//字典序快速排序
func RankLex(values []string) {
	if len(values) <= 1 {
		return
	}
	mid, i := values[0], 1
	head, tail := 0, len(values)-1
	for head < tail {
		if strings.Compare(values[i], mid) == 1 {
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	values[head] = mid
	RankLex(values[:head])
	RankLex(values[head+1:])
}
