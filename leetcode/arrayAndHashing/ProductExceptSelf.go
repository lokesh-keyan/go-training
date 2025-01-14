package leetcode

func productExceptSelf(nums []int) []int {
    n := len(nums)
    result := make([]int, n)
    for i := range result{
        result[i] = 1
    }
    
    for i := 1; i < n; i++{
        result[i] = result[i-1] * nums[i-1]
    }

    rightProduct := 1

    for i := n - 2; i >= 0; i--{
        rightProduct *= nums[i+1]
        result[i] *= rightProduct
    }

    return result
}