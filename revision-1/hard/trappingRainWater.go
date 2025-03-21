package hard

func trap(height []int) int {

	l := 0
	r := len(height) - 1

	lmax := height[0]
	rmax := height[1]
	water := 0

	for l < r {

		if lmax < rmax {
			lmax := max(lmax, height[l])
			water += lmax - height[l]
			l++
		} else {
			rmax := max(rmax, height[r])
			water += rmax - height[r]
			r--
		}

	}
	return water

}
