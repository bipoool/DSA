package medium

func canFinish(numCourses int, prerequisites [][]int) bool {

	graf := map[int][]int{}

	for _, pr := range prerequisites {

		should := pr[1]
		forCourse := pr[0]

		if _, ok := graf[forCourse]; !ok {
			graf[forCourse] = make([]int, 0)
		}
		graf[forCourse] = append(graf[forCourse], should)

	}

	visited := map[int]int{}

	for k := range graf {
		if visited[k] == 2 {
			continue
		}
		if !bfs(graf, visited, k) {
			return false
		}
	}
	return true
}

func bfs(graf map[int][]int, visited map[int]int, currNode int) bool {
	if v, ok := visited[currNode]; ok && v == 1 {
		return false
	}
	if len(graf[currNode]) <= 0 {
		return true
	}
	visited[currNode] = 1
	for _, v := range graf[currNode] {
		if !bfs(graf, visited, v) {
			return false
		}
	}
	visited[currNode] = 2
	return true
}

// func main() {
// 	pre := [][]int{
// 		[]int{1, 4},
// 		[]int{2, 4},
// 		[]int{3, 1},
// 		[]int{3, 2},
// 	}
// 	println(canFinish(2, pre))
// }
