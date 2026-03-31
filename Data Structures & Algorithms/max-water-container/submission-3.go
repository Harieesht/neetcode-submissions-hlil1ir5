func maxArea(heights []int) int {
    left := 0
    right := len(heights)-1
    maxArea := 0

    for {
        if left>=right {
            break
        }
        minHeight := min(heights[left],heights[right])
        area := minHeight * (right-left)
        maxArea = max(maxArea,area)
        if heights[left] < heights[right] {
            left++
        } else {
            right--
        }
    }

    return maxArea

}
