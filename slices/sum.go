package slices

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum = sum + number
	}
	return sum
}

func SumAll(arrOfArr ...[]int) []int {
	result := make([]int, len(arrOfArr))
	for i, arr := range arrOfArr {
		sum := Sum(arr)
		result[i] = sum
	}
	return result
}

func SumAllTails(arrOfArr ...[]int) []int {
	var result []int
	for _, arr := range arrOfArr {
		if len(arr) == 0 {
			result = append(result, 0)
		} else {
			sum := Sum(arr[1:])
			result = append(result, sum)
		}

	}
	return result
}
