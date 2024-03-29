func getOrder(tasks [][]int) []int {
	var st Arr

	for i := 0; i < len(tasks); i++ {
		temp := &Item{
			start:    tasks[i][0],
			duration: tasks[i][1],
			priority: i,
		}
		st.Push(temp)
	}

	sort.Sort(st)
	pq := &PQ{}
	heap.Init(pq)
	heap.Push(pq, st[0])
	curr := st[0].start
	res := make([]int, 0)
	for i := 1; i < len(st); {
		if curr < st[i].start && len(*pq) == 0 {
			heap.Push(pq, st[i])
			curr = st[i].start
			i++
		} else if curr >= st[i].start {
			heap.Push(pq, st[i])
			i++
		} else {
			if len(*pq) > 0 {
				temp := heap.Pop(pq).(*Item)
				fmt.Println("Poping1...")
				fmt.Println(temp)
				res = append(res, temp.priority)
				curr = curr + temp.duration
			} else {
				i++
				curr = 999
			}
		}
	}

	for len(*pq) > 0 {
		temp := heap.Pop(pq).(*Item)
		fmt.Println("Poping2...")
		fmt.Println(temp)
		res = append(res, temp.priority)
	}
	return res
}

type Item struct {
	start    int
	duration int
	index    int
	priority int
}

type PQ []*Item

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Less(i, j int) bool {
	return pq[i].duration < pq[j].duration || (pq[i].duration == pq[j].duration && pq[i].priority < pq[j].priority)
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

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

type Arr []*Item

func (pq Arr) Len() int { return len(pq) }
func (pq Arr) Less(i, j int) bool {
	return pq[i].start < pq[j].start
}
func (pq Arr) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = pq[j].index
	pq[j].index = pq[i].index
}

func (pq *Arr) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
