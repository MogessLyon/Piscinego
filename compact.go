package piscine

func Compact(ptr *[]string) int {
	var counter int
	var arr []string
	for _, value := range *ptr {
		if value != "" {
			arr = append(arr, value)
			counter++
		}
	}

	*ptr = arr
	return counter
}
