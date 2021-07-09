package insertsort

func InsertionSort(A []int) []int {
	slice := make([]int, len(A))
	copy(slice, A)
	for i := 1; i < len(slice); i++ {
		key := slice[i]
		lst := i
		for j := i - 1; j > -1; j-- {
			if slice[j] < key {
				break
			}
			slice[j+1] = slice[j]
			lst = j
		}
		slice[lst] = key
	}
	return slice
}
