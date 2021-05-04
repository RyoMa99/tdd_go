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

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, nums := range numbersToSum {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			tails := nums[1:]
			sums = append(sums, Sum(tails))
		}
	}
	return sums
}
