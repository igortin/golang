package code

import "fmt"

func main() {

	nums := []int{0,0,1,1,1,2,2,3,3,4} //
	fmt.Println(removeDuplicates(nums))
}


func removeDuplicates(nums []int) int {
	i := 0
	if len(nums) == 0 {
		return len(nums)
	}
	maxval := nums[len(nums)-1]

	for j := i + 1; j < len(nums); j++ {
		if nums[i] < nums[j] {
			i++
		} else {
			nums[j] = get_value(nums, i)
			i++
		}
		if nums[j] == maxval {
			nums = nums[:j+1]
			break
			} else if nums[j] > maxval {
				nums = nums[:j]
		}

	}
	fmt.Println(nums)
	return len(nums)
}
func get_value(nums []int, i int) int {
	for j := i + 1 ; j < len(nums); j++ {
		if nums[i] < nums[j] {
			fmt.Println(nums[j])
			return nums[j]
		}
	}
	return nums[i]+1
}



