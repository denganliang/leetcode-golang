func BiarySearch(nums []int, target int, low int, high int) int {
    if high < low {
        return -1
    }
    mid := low + (high - low)/2
    if nums[mid] == target {
        return mid
    }
    left := BiarySearch(nums, target, low, mid-1)
    if left != -1 {
        return left
    }
    right := BiarySearch(nums, target, mid+1, high)
    if right != -1 {
        return right
    }
    return -1
}
func search(nums []int, target int) int {
    return BiarySearch(nums, target, 0, len(nums)-1)
}
