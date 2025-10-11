package algo

func Sort(num []int) []int {

	for i := 1; i < len(num); i++ {
		key := num[i]
		j := i - 1
		for j > -1 && key < num[j] {
			swap(&num[j], &num[j+1])
			j--
		}
	}

	return num
}

func swap(a, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}
