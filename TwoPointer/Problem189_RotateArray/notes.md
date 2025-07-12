the ideal solution is when you visualize the order in which the problem is done
1. find the k module if k > len(nums)
    k = k % len(nums)
2. reverse the nums
3. reverse the nums upto k - 1
4. reverse the nums from k to len(nums) - 1