package tools

import (
	"encoding/csv"
	"log"
	"os"
)

func errHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Dataset struct {
	Values [][]string
	Size   float64
	Len    int
}

//用来加载csv文件的工具
func Load(file string) *Dataset {
	dataset, err := os.Open(file)
	errHandler(err)
	defer dataset.Close()

	r := csv.NewReader(dataset)
	records, err := r.ReadAll()
	errHandler(err)

	return &Dataset{
		Values: records,
		Size:   (float64)(len(records)),
		Len:    len(records[0]),
	}
}

//对首项为标头的数据
//逐行字母序清理
func Clean(dataset *Dataset) {
	for i, record := range dataset.Values {
		dataset.Values[i] = record[1:]
		RankLex(dataset.Values[i])
	}
}

func CleanBuyFre(dataset *Dataset, mark map[string]float64) {
	for i, record := range dataset.Values {
		dataset.Values[i] = record[1:]
		RankFre(dataset.Values[i], mark, 1)
	}
}

//用来写入csv文件的工具,目前假设是写入map
func Store() {}
