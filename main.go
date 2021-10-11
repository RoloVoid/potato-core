package main

import (
	"core/methods"
	"core/tools"
	"fmt"
)

func main() {

	fmt.Println("主程序入口")
	records := tools.Load("./dataset/shopping.csv")
	tools.Clean(records)

	methods.InitializeItems(records)
}
