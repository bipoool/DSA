package medium

// Add node and traverse till you find cycle
// Make sure to traverse whenever you add a cycle
var vis map[int]int

func findRedundantConnection(edges [][]int) []int {

	vis = map[int]int{}
	m := map[int][]int{}

	for i := range edges {

		a := edges[i][0]
		b := edges[i][1]

		if m[a] == nil {
			m[a] = []int{}
		}

		if m[b] == nil {
			m[b] = []int{}
		}

		m[a] = append(m[a], b)
		m[b] = append(m[b], a)

		if help(m, a, 0) {
			return []int{a, b}
		}

	}
	return nil

}

func help(m map[int][]int, n int, p int) bool {

	if m[n] == nil || vis[n] == 1 {
		return true
	}

	vis[n] = 1
	for i := range m[n] {
		if m[n][i] == p {
			continue
		}
		if help(m, m[n][i], n) {
			return true
		}
	}
	vis[n] = 0
	return false
}
