func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

    var ans []int
    
    for i := 0; i < len(nums); i++ {
        diff := target - nums[i]
        if j, ok := m[diff]; ok {
            ans = append(ans, j)
            ans = append(ans, i)
            return ans
        }
        m[nums[i]] = i
    }

    return ans
}
