type Node struct {
	in int
	child []int
}

func canFinish(numCourses int, prerequisites [][]int) bool {
    pre := prerequisites
	g := make([]*Node, 0)
	for i:=0; i < numCourses; i++ {
		g = append(g, &Node{
			0,
			make([]int, 0),
		})
	}
    for i := 0; i < len(pre); i++ {
		g[pre[i][0]].in++
		g[pre[i][1]].child = append(g[pre[i][1]].child, pre[i][0])
	}
	res := make([]int, 0)
	for i := 0; i < len(g); i++ {
		if g[i].in == 0 {
			res = append(res, i)
		}
	}

	if len(res) == 0 {
		return false
	}

	for len(res) > 0 {
		temp := res[0]
		n := len(res)
		res = res[1:n]
		for i := 0; i < len(g[temp].child); i++ {
			ch := g[temp].child[i]
			g[ch].in--
			if g[ch].in == 0 {
				res = append(res, ch)
			}
		}
	}

	for i := 0; i < len(g); i++ {
		if g[i].in > 0 {
			return false
		}
	}
	return true
}