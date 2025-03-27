package medium

func searchMatrix(matrix [][]int, target int) bool {

	l := 0
	r := len(matrix) - 1
	targetRow := -1

	for l <= r {

		mid := (l + r) / 2

		arr := matrix[mid]

		if arr[0] <= target && arr[len(arr)-1] >= target {
			targetRow = mid
			break
		}

		if arr[0] > target {
			r = mid + 1
		} else {
			l = mid - 1
		}

	}

	if targetRow == -1 {
		return false
	}

	l = 0
	r = len(matrix[targetRow]) - 1

	for l <= r {

		mid := (l + r) / 2

		if matrix[targetRow][mid] == target {
			return true
		}

		if matrix[targetRow][mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}

	}
	return false
}
