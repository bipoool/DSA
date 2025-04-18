package medium

// Sliding window
// Just check the max frequency element + k should not be less than the currLen

func characterReplacement(s string, k int) int {

	l := 0
	r := 0
	mxc := 0

	hmap := map[byte]int{}
	mxlen := 0

	for r < len(s) {

		hmap[s[r]]++
		if mxc < hmap[s[r]] {
			mxc = hmap[s[r]]
		}
		for (r-l+1)-mxc > k {
			hmap[s[l]]--
			l++
		}
		mxlen = max(mxlen, r-l+1)
		r++

	}
	return mxlen

}
