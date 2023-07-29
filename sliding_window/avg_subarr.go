package sliding_window

/*
Average Of All Subarrays - Example Question
Given an array, find the average of all subarrays of ‘K’ contiguous elements in it.

Let’s understand this problem with real input:

Array: [1, 3, 2, 6, -1, 4, 1, 8, 2], K=5
Here is the final output containing the averages of all subarrays of size '5':

Output: [2.2, 2.8, 2.4, 3.6, 2.8]
*/

func FindAvgSubArray(numbers []int, k int) []float32 {
	var ans []float32
	sum := 0
	var avg float32
	for i := 0; i < k; i++ {
		sum += numbers[i]
	}
	avg = float32(sum / k)
	ans = append(ans, avg)

	left := 0
	for i := k; i < len(numbers); i++ {
		sum = sum + numbers[i] - numbers[left]
		avg = float32(sum / k)
		ans = append(ans, avg)
		left += 1
	}

	return ans
}
