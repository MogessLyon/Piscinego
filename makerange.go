package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	arr := make([]int, max-min)

	value := min
	for i := 0; i < max-min; i++ {
		arr[i] = value
		value += 1
	}

	return arr
}
