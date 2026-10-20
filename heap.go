package main

//heap

type pqi struct{ x int64 }
type priorityQueue struct{ a []pqi }

func (a priorityQueue) Len() int           { return len(a.a) }
func (a priorityQueue) Swap(i, j int)      { a.a[i], a.a[j] = a.a[j], a.a[i] }
func (a priorityQueue) Less(i, j int) bool { return a.a[i].x > a.a[j].x }
func (a *priorityQueue) Push(x any) {
	y, _ := x.(pqi)
	a.a = append(a.a, y)
}
func (a *priorityQueue) Pop() any {
	b := a.a[len(a.a)-1]
	a.a = a.a[:len(a.a)-1]
	return b
}
