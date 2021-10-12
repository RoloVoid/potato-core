package methods

import (
	"core/tools"
	"fmt"
	"log"
	"os"
	"strconv"
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
	endfix     string
	dataset    *tools.Dataset
)

//init 过程中读取数据集，此处先不实现
func init() {
	min_sup = 0.002
	min_conf = 0.7
	itemSet = make(map[string]float64)
	datasetDir = "./dataset"
	targetDir = datasetDir
	dataset = generateDatabase()
	endfix = ".csv"
}

func errHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func generateDatabase() *tools.Dataset {
	records := tools.Load("./dataset/shopping.csv")
	tools.Clean(records)
	return records
}

//生成一项集，计算支持度，保留三位小数
//一项集格式：项,支持度,频度
func InitializeItems() {
	file, err := os.OpenFile("./dataset/1.csv", os.O_WRONLY|os.O_CREATE, 0666)
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
func One2Two() {
	file, err := os.OpenFile("./dataset/2.csv", os.O_WRONLY|os.O_CREATE, 0666)
	errHandler(err)
	defer file.Close()

	var records, models [][]string
	oneitemset := tools.Load("./dataset/1.csv")
	records = oneitemset.Values
	models = dataset.Values
	for i := 0; i < len(records); i++ {
		for j := i + 1; j < len(records); j++ {
			counter := 0.0
			for _, data := range models {
				model := strings.Join(data, ",")
				if strings.Contains(model, records[i][0]) && strings.Contains(model, records[j][0]) {
					counter++
				}
			}
			if sup := counter / dataset.Size; sup > min_sup {
				fmt.Fprintln(file, records[i][0]+","+records[j][0]+","+fmt.Sprintf("%.3f", sup)+","+fmt.Sprintf("%.f", counter))
			}
		}
	}
}

//二项集以上，基于k项集生成k+1项集,k指项数
func Next(k int) {
	path := targetDir + "/" + strconv.Itoa(k) + endfix
	path2 := targetDir + "/" + strconv.Itoa(k+1) + endfix
	kset := tools.Load(path)

	file, err := os.OpenFile(path2, os.O_WRONLY|os.O_CREATE, 0666)
	errHandler(err)
	defer file.Close()

	raw := dataset.Values
	krecord := kset.Values
	for i := 0; i < len(krecord); i++ {
		for j := i + 1; j < len(krecord); j++ {
			counter := 0.0
			a := strings.Join(krecord[i][0:(k-1)], ",")
			b := strings.Join(krecord[j], ",")
			if strings.Contains(b, a) {
				for _, models := range raw {
					model := strings.Join(models, ",")
					if !strings.Contains(model, krecord[j][k-1]) {
						break
					}
					for _, item := range krecord[i] {
						if !strings.Contains(model, item) {
							break
						}
					}
					counter++
				}
			}
			if sup := counter / dataset.Size; sup > min_sup {
				fmt.Fprintln(file, a+","+krecord[i][k-1]+","+krecord[j][k-1]+","+fmt.Sprintf("%.3f", sup)+","+fmt.Sprintf("%.f", counter))
			}
		}
	}
}
