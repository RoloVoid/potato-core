package apriori

import (
	"core/tools"
	"fmt"
	"os"
	"strings"
)

type Apriori int

//以下是传统的apriori实现
var (
	min_sup_rate  float64
	min_sup       float64
	min_conf_rate float64
	items         []string
	itemSet       map[string]float64
	datasetDir    string
	targetDir     string
	endfix        string
	dataset       *tools.Dataset
)

//fg树相关的属性
var (
	fgcopy   freHash = itemSet
	linklist linkHash
)

//init 过程中读取数据集，此处先不实现
func init() {
	min_sup_rate = 0.002
	min_conf_rate = 0.7
	min_sup = 5
	itemSet = make(map[string]float64)
	datasetDir = "./dataset"
	targetDir = datasetDir
	dataset = generateDatabase()
	endfix = ".csv"
	linklist = make(map[string]*TreeNode)
	items = make([]string, 30)
}

//生成一项集，可计算支持度，保留三位小数
//一项集格式：项,支持度,频度
//choose用于声明是否写入文件,0为不写入，1为写入
func InitializeItems(choose int) {
	for _, records := range dataset.Values {
		for _, item := range records {
			item = strings.TrimSpace(item)
			if item == "" || item == " " {
				continue
			}
			if value, ok := itemSet[item]; ok {
				itemSet[item] = value + 1
			} else {
				itemSet[item] = 1
			}
		}
	}
	//写入文件的处理
	if choose == 1 {
		file, err := os.OpenFile("./dataset/1.csv", os.O_WRONLY|os.O_CREATE, 0666)
		errHandler(err)
		defer file.Close()
		for key, value := range itemSet {
			if measure(value, 1) {
				items = append(items, key)
				fmt.Fprintln(file, key+","+fmt.Sprintf("%.f", value))
			}
		}
	}
	for key, value := range itemSet {
		if measure(value, 1) {
			items = append(items, key)
		}
	}
	tools.RankFre(items, itemSet, 1)
}
