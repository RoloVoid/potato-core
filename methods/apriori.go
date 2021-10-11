package methods

type vector struct {
	support    float64 //AB同时出现的概率
	confidence float64 //若A则B的条件概率
}

type linkedlist struct {
	next  *linkedlist
	value string
}

type Item struct {
	linkedlist
	vector
	length int
}

var (
	min_sup  float64
	min_conf float64
	itemSet  map[*Item]*linkedlist
)

func init() {
	min_sup = 0.5
	min_conf = 0.7
	itemSet = make(map[*Item]*linkedlist)
}

func (item *Item) Search() {

}
