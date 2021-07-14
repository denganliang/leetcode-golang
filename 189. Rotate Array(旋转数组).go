func rotate(nums []int, k int)  {
    n := k%len(nums)
    copy(nums, append(nums[len(nums)-n:], nums[:len(nums)-n]...))
}
