package main

func Sum(numbers []int) (ans int) {
	for _, num := range numbers {
		ans += num
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, nums := range numbersToSum {
		sums = append(sums, Sum(nums))
	}
	return sums
}
