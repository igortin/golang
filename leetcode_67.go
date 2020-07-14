func maxSubArray(nums []int) int {
	var sum = 0
	var MaxInt  = int(^uint(0) >> 1)
	var MinInt  = -MaxInt - 1
	var maxSuma = MinInt
	
	for _, v := range nums {
		sum = sum + v
		if sum < 0 {
			sum = 0
		}
		maxSuma = max(sum,maxSuma)
	}
	
    if maxSuma == 0 {
        var no = nums[0] 
        for _, v := range nums{
            no=max(no,v)
        }
    return no
   }
return maxSuma
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
