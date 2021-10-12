package main

import (
	"core/methods"
	"fmt"
)

func main() {

	fmt.Println("主程序入口")
	// records := tools.Load("./dataset/shopping.csv")
	// tools.Clean(records)

	methods.InitializeItems()
	methods.One2Two()
	methods.Next(2)
	methods.Next(3)
	methods.Next(4)
	methods.Next(5)
}
