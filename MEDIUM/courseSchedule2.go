package medium

var fres []int
var visitedCourse []int

func findOrder(numCourses int, prerequisites [][]int) []int {

	hmap := map[int][]int{}
	fres = make([]int, numCourses)

	for _, pre := range prerequisites {
		c1 := pre[0]
		c2 := pre[1]
		if _, ok := hmap[c1]; !ok {
			hmap[c1] = make([]int, 0)
		}
		hmap[c1] = append(hmap[c1], c2)
	}
	return fres
}

func helperOrder(num int, hmap map[int][]int, curr int, res []int) {

	if len(res) == num {
		copy(fres, res)
		return
	}

	visitedCourse[curr] = 1
	for _, c := range hmap[curr] {
		if visitedCourse[c] == 1 {
			continue
		}
		res = append(res, c)
		helperOrder(num, hmap, c, res)
		res = res[:len(res)-1]
	}
	visitedCourse[curr] = 0

}
