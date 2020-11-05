package leetcode

func nonrepeat(s string )int{
	lastOccurred := make(map[rune]int)

	start := 0

	maxLength := 0

	for i, ch := range []rune(s) {

		if last_i, ok := lastOccurred[ch]; ok && last_i >= start {
			start = last_i + 1
			//start = last_i
		}

		if i-start+1 > maxLength {
			//if i-start > maxLength {
			//	maxLength = i - start
			maxLength = i - start + 1
		}

		lastOccurred[ch] = i

		//a i=0  map a = 0  max =1
		//b i=1  map b = 1 max =2
		//c i=2 map  c = 2 max = 3

		//a  i=3 start = 1  start =1
	}

	return maxLength

}
