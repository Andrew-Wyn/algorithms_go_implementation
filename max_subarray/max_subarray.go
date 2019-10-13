
/*

	MaxSubArray GO implementation

	complessita computazionale del max subarray tramite divide-et-impera

	T(n) = {
		O(1), n=1
		2T(n/2) + O(n), n>1
	}

	applicando il teorema dell'esperto abbiamo che O(T(n)) è uguale a nlogn

*/


package main

import (
	"fmt"
)

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

func find_max_crossing_subarray(v []int, low int, mid int, high int) (int, int, int) {
	fmt.Println(MinInt)
	var left_sum int = MinInt
	var sum int = 0

	var max_left int = 0

	for i := mid; i>=0; i-- {
		sum += v[i]
		if sum > left_sum {
			left_sum = sum
			max_left = i
		}
	}

	var right_sum int = MinInt
	sum = 0

	var max_right int = 0

	for i := mid+1; i<=high; i++ {
		sum += v[i]
		if sum > right_sum {
			right_sum = sum
			max_right = i
		}
	}

	return max_left, max_right, left_sum+right_sum
}

func find_max_subarray(v []int, low int, high int) (int, int, int) {
	if high == low {
		return low, high, v[low]
	} else {
		var mid int = (low+high)/2
		left_low, left_high, left_sum := find_max_subarray(v, low, mid)
		right_low, right_high, right_sum := find_max_subarray(v, mid+1, high)
		cross_low, cross_high, cross_sum := find_max_crossing_subarray(v, low, mid, high)
		if left_sum >= right_sum && left_sum >= cross_sum {
			return left_low, left_high, left_sum
		} else if right_sum >= left_sum && right_sum >= cross_sum {
			return right_low, right_high, right_sum
		} else {
			return cross_low, cross_high, cross_sum
		}
	}
}

func main() {
	fmt.Println("Max SubArray algorithm")

	var vector = []int {1,2,-1, 2}

	fmt.Println(vector)
	fmt.Println(find_max_subarray(vector, 0, len(vector) - 1))

}
