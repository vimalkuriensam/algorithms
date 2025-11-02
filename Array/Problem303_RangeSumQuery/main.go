package main

import "fmt"

type NumArray struct {
	prefix []int
}

func Constructor(nums []int) NumArray {
	prefix := make([]int, len(nums)+1)
	prefix[0] = 0
	for index, num := range nums {
		prefix[index+1] = prefix[index] + num
	}
	return NumArray{prefix: prefix}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefix[right+1] - this.prefix[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

func main() {
	numArray := Constructor([]int{-2, 0, 3, -5, 2, -1})
	ans1 := numArray.SumRange(0, 2)
	ans2 := numArray.SumRange(2, 5)
	ans3 := numArray.SumRange(0, 5)
	fmt.Println(ans1, ans2, ans3)
}
