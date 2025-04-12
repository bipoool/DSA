package medium

type DetectSquares struct {
	mapall map[int]map[int]int
}

func ConstructorDetectSquares() DetectSquares {
	return DetectSquares{
		mapall: map[int]map[int]int{},
	}
}

func (this *DetectSquares) Add(point []int) {
	if _, ok := this.mapall[point[0]]; !ok {
		this.mapall[point[0]] = map[int]int{}
	}
	this.mapall[point[0]][point[1]]++
}

func (this *DetectSquares) Count(point []int) int {

	count := 0

	x1 := point[0]
	y1 := point[1]

	ally, ok := this.mapall[x1]
	if !ok {
		return count
	}
	for y := range ally {

		dist := y - y1
		if dist == 0 {
			continue
		}
		x2 := x1 - dist
		x3 := x1 + dist

		if _, ok := this.mapall[x2]; ok {
			count += this.mapall[x1][y] * this.mapall[x2][y] * this.mapall[x2][y1]
		}

		if _, ok := this.mapall[x3]; ok {
			count += this.mapall[x1][y] * this.mapall[x3][y] * this.mapall[x3][y1]
		}
	}

	return count
}
