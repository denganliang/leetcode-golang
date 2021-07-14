/*
//simple solution, two pointer
func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
func sortedSquares(nums []int) []int {
    var ret []int
    start := 0
    end := len(nums)-1
    for start < end {
        if abs(nums[start]) > abs(nums[end]) {
            ret = append([]int{nums[start]*nums[start]}, ret...)
            start++
        } else {
            ret = append([]int{nums[end]*nums[end]}, ret...)
            end--
        }
    }
    ret = append([]int{nums[start]*nums[start]}, ret...)
    return ret
}

*/

//insert sort and output
func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
func BiarySearch(nums []int, left int, right int) int {
    if left >= right {
        if nums[left] < 0 {
            return -1
        }
        return left
    }
    mid := left + (right-left)/2
    if mid > 0 {
        if nums[mid]>=0  {
            if nums[mid-1]<0 {
                return mid  
            } else {
                return BiarySearch(nums,left, mid-1)
            }
        } else {
            return BiarySearch(nums, mid+1, right)
        }
    } else {
        for left <= right {
            if nums[left]>=0 {
                return left
            }
            left++
        }
        return -1
    }
}
func sortedSquares(nums []int) []int {
    idx := BiarySearch(nums, 0, len(nums)-1)
    if idx == -1 {
        idx = len(nums)
    }
    left := nums[:idx]
    right := nums[idx:]
    if len(left) > 0 {
        idx = 0
        il := len(left)-1
        for il >= 0 {
            for idx<len(right) && right[idx] <= abs(left[il]) {
                idx++
            }
            if idx == len(right) {
                right = append(right, left[il])
            } else {
                right = append(right[:idx+1], right[idx:]...)
                right[idx] = left[il]
            }
            il--
            idx++
        }
    }
    var ret []int 
    for _,v:=range right {
        ret = append(ret, v*v)
    }
    return ret
}
