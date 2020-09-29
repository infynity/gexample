package lcd
func twoSum(nums []int, target int) []int {
	ma:=make(map[int]int)
	for i,v:=range nums{

		desum := target-v
		if value,exist:=ma[desum];exist{

			return []int{i,value}
		}

		ma[v]=i
	}


	return nil
}