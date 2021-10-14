package main

import (
	ap "core/methods/apriori"
	"fmt"
)

func main() {

	fmt.Println("主程序入口")
	// records := tools.Load("./dataset/shopping.csv")
	// tools.Clean(records)
	// methods.One2Two()
	// methods.Next(2)
	// methods.Next(3)
	// methods.Next(4)
	// methods.Next(5)
	// test := make([]int, 0, 5)
	// fmt.Println(append(test, 8))
	var a ap.Apriori
	a.FPTree()
}
