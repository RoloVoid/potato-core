package apriori

import (
	"core/tools"
	"fmt"
	"os"
	"strings"
)

type linkHash map[string]*TreeNode
type freHash map[string]float64

type TreeNode struct {
	fre   float64
	value string
	plink *TreeNode
	nlink *TreeNode
	child linkHash
}

func initTreeNode() *TreeNode {
	node := new(TreeNode)
	node.child = make(map[string]*TreeNode)
	return node
}

//生成的时候从上往下，挖掘的时候从下往上
//此处空间换时间，用hash表维护子节点，若数据量过大，建议使用广义表
func generateTree(start *TreeNode) {
	InitializeItems(1)
	for _, records := range dataset.Values {
		tools.RankFre(records, itemSet, 0)
		root := start
		for _, item := range records {
			if item == "" || item == " " {
				continue
			}
			if node, ok := root.child[item]; ok {
				temp := node.fre
				node.fre = temp + 1
				root = node
			} else {
				new := initTreeNode()
				new.value = item
				new.fre = 1
				new.plink = root

				//串接link,没有就新建一个记录
				if data := linklist[item]; data != nil {
					new.nlink = data
				}
				linklist[item] = new
				root.child[item] = new
				root = new
			}
		}
	}
}

func getFre() {
	file, err := os.OpenFile("./dataset/result.csv", os.O_WRONLY|os.O_CREATE, 0666)
	errHandler(err)
	defer file.Close()
	//
	fmt.Println(items)
	//
	for i, item := range items {
		//
		if i == 0 || i == 1 || i == 2 {
			fmt.Println(item)
		}
		//
		if item == "" || item == " " {
			continue
		}
		link := linklist[item]
		if link == nil || link.value == "" {
			continue
		}
		for link.nlink != nil {
			if link.fre < min_sup {
				link = link.nlink
				continue
			}
			p := link
			list := make([]string, 0)
			for p.plink != nil {
				list = append(list, p.value)
				p = p.plink
			}
			data := strings.Join(list, ",")
			fmt.Fprintln(file, data+"/"+fmt.Sprintf("%.f", link.fre))
			link = link.nlink
		}
	}
}

//暴露接口
func (ap *Apriori) getRoot() *TreeNode {
	return initTreeNode()
}

func (ap *Apriori) FPTree() {
	root := ap.getRoot()

	generateTree(root)
	fmt.Println(2)
	getFre()
	fmt.Println(3)
	// fmt.Println(root)
}
