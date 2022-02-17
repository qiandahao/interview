package main

import "fmt"

func main() {
	tasks := [][]int{
		{7, 10},
		{7, 12},
		{7, 5},
		{7, 2},
	}
	fmt.Println(tasks)
}

type Item struct {
	start int
	duration int
	index int
	priority int
}

type PQ []*Item

func (pq PQ) Len() int {return len(pq)}
func (pq PQ) Less(i, j int) bool {
	return pq[i].duration > pq[j].duration
}
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = pq[j].index
	pq[j].index = pq[i].index
}

func (pq *PQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PQ) Pop() interface{}{
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.
}

