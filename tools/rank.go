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

//基于频度来排序,1为从小到大，0为从大到小
func RankFre(values []string, mark map[string]float64, choose int) {
	if len(values) <= 1 {
		return
	}
	mid, i := strings.TrimSpace(values[0]), 1
	head, tail := 0, len(values)-1
	for head < tail {
		//从小到大
		temp1 := strings.TrimSpace(values[i])
		if mark[temp1] > mark[mid] && choose == 1 {
			values[i], values[tail] = values[tail], values[i]
			tail--
		}
		//从大到小
		if mark[temp1] < mark[mid] && choose == 0 {
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	values[head] = mid
	RankFre(values[:head], mark, choose)
	RankFre(values[head+1:], mark, choose)
}
