func BiarySearch(nums []int, target int, left int, right int) int {
    if left > right {
        return left
    }
    mid := left + (right-left)/2
    if nums[mid] == target {
        return mid
    } else if nums[mid] > target {
        return BiarySearch(nums, target, left, mid-1)
    } 
    return BiarySearch(nums, target, mid+1, right)
}

func searchInsert(nums []int, target int) int {
    return BiarySearch(nums, target, 0, len(nums)-1)
}
