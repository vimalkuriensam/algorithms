package controllers

// Merge Sort Algorithm
// Sort(arr1, arr2)
// 		arr3 = arr[arr1.length+arr2.length]
// 		for i < arr1.length && j < arr2.length
// 			if arr1[i] < arr2[j]
// 				arr3 = arr1[i]
// 				i++
// 			else
// 				arr3 = arr2[j]
// 				j++
// 		arr3 = [arr3, ...arr1, ...arr2]
// 		return arr3
//
// Merge(Arr[N])
// 		middle = Length(N)
// 		left = Arr[middle:]
//		right = Arr[:middle]
//		leftSort = Merge(left)
//		rightSort = Merge(right)
//		Sort(left, right)

func (ms *Adaptor) SortTwoArrays(left []int, right []int) []int {
	i := 0
	j := 0
	k := 0
	arr3 := make([]int, len(left)+len(right))
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			arr3[k] = left[i]
			i++
		} else {
			arr3[k] = right[j]
			j++
		}
		k++
	}
	if i < len(left) {
		for i < len(left) {
			arr3[k] = left[i]
			i, k = i+1, k+1
		}
	}
	if j < len(right) {
		for j < len(right) {
			arr3[k] = right[j]
			j, k = j+1, k+1
		}
		arr3 = append(arr3, right[j:]...)
	}
	return arr3
}

func (ms *Adaptor) MergeSort(nums []int) []int {
	if len(nums) > 1 {
		length := len(nums)
		left := nums[:length/2]
		right := nums[length/2:]
		leftRec := ms.MergeSort(left) // [5, 2, 7, 1]
		rightRec := ms.MergeSort(right)
		nums = ms.SortTwoArrays(leftRec, rightRec)
	}
	return nums
}
