package hard

import "fmt"

func candy(r []int) int {

	ans := make([]int, len(r))
	ans[0] = 1
	for i := 1; i < len(r); i++ {
		ans[i] = 1
		if r[i] > r[i-1] {
			ans[i] = ans[i-1] + 1
		}
	}
	fmt.Println(ans)
	for i := len(r) - 2; i >= 0; i-- {
		if r[i] > r[i+1] && ans[i] <= ans[i+1] {
			ans[i] = ans[i+1] + 1
		}
	}
	sum := 0

	for i := range ans {
		sum += ans[i]
	}
	return sum

}
