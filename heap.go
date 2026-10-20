import "container/heap"

type pqi struct{ x int64 }
type priorityQueue []pqi

func (a priorityQueue) Len() int           { return len(a) }
func (a priorityQueue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a priorityQueue) Less(i, j int) bool { return a[i].x > a[j].x }
func (a *priorityQueue) Push(x any) {
	y, _ := x.(pqi)
	*a = append(*a, y)
}
func (a *priorityQueue) Pop() any {
	b := (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	return b
}
