package arrays

func Sum(numbers []int) (sum int) {

	for _, number := range numbers { // using the blank identifier to ignore the index value as we dont' need it
		sum += number
	}

	return
}

func SumAll(list ...[]int) (sum []int) {

	for _, numbers := range list {
		result := Sum(numbers)
		sum = append(sum, result)
	}
	return
}

func SumAllTails(tosum ...[]int) (sums []int) {

	for _, list := range tosum {
		if len(list) == 0 {
			sums = append(sums, 0)
		} else {
			tails := list[1:] // slices can be sliced and here we are slicing them by providing the low to high range and if you omit one side then it will capture to the end of that side.
			sums = append(sums, Sum(tails))
		}
	}
	return
}
