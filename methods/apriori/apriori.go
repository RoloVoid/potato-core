package apriori

import (
	"core/tools"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func measure(fre float64, choose int) bool {
	if choose == 0 {
		return fre/dataset.Size > min_sup_rate
	} else {
		return fre > min_sup
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
			if measure(counter, 1) {
				fmt.Fprintln(file, records[i][0]+","+records[j][0]+","+fmt.Sprintf("%.f", counter))
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
			if measure(counter, 1) {
				fmt.Fprintln(file, a+","+krecord[i][k-1]+","+krecord[j][k-1]+","+fmt.Sprintf("%.f", counter))
			}
		}
	}
}
