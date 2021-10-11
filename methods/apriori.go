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
	min_sup    float64
	min_conf   float64
	itemSet    map[string]float64
	datasetDir string
	targetDir  string
	dataset *tools.Dataset
)

//init 过程中读取数据集，此处先不实现
func init() {
	min_sup = 0.01
	min_conf = 0.7
	itemSet = make(map[string]float64)
	datasetDir = "./dataset"
	dataset := generateDatabase()
}

func errHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func generateDatabase() *tools.Dataset{
	records := tools.Load("./dataset/shopping.csv")
	tools.Clean(records)
	return records
}

//生成一项集，计算支持度，保留三位小数
//一项集格式：项,支持度,频度
func InitializeItems(dataset *tools.Dataset) {
	file, err := os.OpenFile("./dataset/1item.csv", os.O_WRONLY|os.O_CREATE, 0666)
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
			fmt.Fprintln(file, key+","+fmt.Sprintf("%.3f", sup)+","+fmt.Sprintf("%.f", value))
		}
	}
}

//一项集生成二项集并剪枝
func One2Two(oneitemfile string) {
	oneitemset := tools.Load("./dataset/1item.csv")
	for _, record := range oneitemset.Values {
		for _, record1 := range oneitemset.Values {
			if record[0] == record1[0] {
				continue
			}
			for 
		}
	}
}
