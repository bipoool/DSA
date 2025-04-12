package medium

func multiply(num1 string, num2 string) string {

	res := make([]int, len(num1)+len(num2)+1)

	for i := len(num1); i >= 0; i-- {
		pos := i
		for j := len(num2); j >= 0; j-- {
			d1 := num1[i] - '0'
			d2 := num2[j] - '0'
			mult := d1 * d2
			res[pos] += int(mult)
			res[pos+1] = res[pos] / 10
			res[pos] = res[pos] % 10
		}
	}

	return ""

}
