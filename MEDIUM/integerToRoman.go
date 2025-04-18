package medium

func intToRoman(num int) string {

	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	rom := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	res := ""
	i := 0
	for num > 0 {
		if num >= nums[i] {
			res += rom[i]
			num -= nums[i]
		} else {
			i++
		}
	}

	return res
}
