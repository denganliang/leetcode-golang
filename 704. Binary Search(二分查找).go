func BiarySearch(nums []int, target int, low int, high int) int {
    if high < low {
        return -1
    }
    mid := low + (high - low)/2
    if nums[mid] == target {
        return mid
    } else if nums[mid] < target {
        return BiarySearch(nums, target, mid+1, high)
    } else {
        return BiarySearch(nums, target, low, mid-1)
    }
}
func search(nums []int, target int) int {
    return BiarySearch(nums, target, 0, len(nums)-1)
}
