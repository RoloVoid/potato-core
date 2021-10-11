package methods

import (
	"core/tools"
	"fmt"
	"log"
	"os"
	"strings"
)

type vector struct {
	support    float64 //AB同时出现的概率
	confidence float64 //若A则B的条件概率
}

type linkedlist struct {
	next   *linkedlist
	lvalue string
}

type Item struct {
	vector
	nextItem *Item
	value    string
	length   int
}

var (
	min_sup  float64
	min_conf float64
	itemSet  map[string]float64
)

func init() {
	min_sup = 0.01
	min_conf = 0.7
	itemSet = make(map[string]float64)
}
func errHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//生成一项集，计算支持度，保留三位小数
func InitializeItems(dataset *tools.Dataset) {
	file, err := os.OpenFile("./dataset/oneitem.csv", os.O_WRONLY|os.O_CREATE, 0666)
	errHandler(err)
	defer file.Close()
	for _, records := range dataset.Values {
		for _, item := range records {
			if item == "" || item == " " {
				continue
			}
			item = strings.TrimSpace(item)
			if value, ok := itemSet[item]; ok {
				itemSet[item] = value + 1
			} else {
				itemSet[item] = 1
			}
		}
	}
	for key, value := range itemSet {
		if sup := value / (dataset.Size); sup > min_sup {
			fmt.Fprintln(file, key+","+fmt.Sprintf("%.3f", sup))
		}
	}
}
