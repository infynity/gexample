package main

import "fmt"

func GetLonest(s string) (n int) {

	var mp = make( map[string]int,0)
	runes := []rune(s)


	var start,maxlenth = 0,0
	for i,v :=range(runes){
		mpKey := string(v)
		if last_idx,exist:=mp[mpKey];exist{
			start = last_idx + 1

			//start = i - last_idx +1
		}

		if i-start+1 > maxlenth {
			//if i-start > maxLength {
			//	maxLength = i - start
			maxlenth = i - start + 1
		}
		//if maxlenth<start{
		//	maxlenth=start
		//}

		mp[mpKey]=i
		fmt.Println(i,v)
	}

	fmt.Println(mp,maxlenth)
	return maxlenth
}

