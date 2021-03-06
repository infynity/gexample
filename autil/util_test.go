package utils

import (
	"fmt"
	"testing"
)

/*func Test_Problem104(t *testing.T) {

	qs := []question104{

		{
			para104{[]int{}},
			ans104{0},
		},

		{
			para104{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
			ans104{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 104------------------------\n")

	for _, q := range qs {
		_, p := q.ans104, q.para104
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", maxDepth(root))
	}
	fmt.Printf("\n\n\n")
}
*/

func TestHuiwen(t *testing.T){
	var a lgd

	 //a.isPalindrome("A man, a plan, a canal: Panama")
	 a.isPalindrome("")
}

type vt struct {
	param *TreeNode
	res  int
}

func TestTmd(t *testing.T){

	tests := []vt{
		{buildTr(), 3},
		{buildTr2(), 1},
		{buildTr3(), 10},
	}

	for _, tt := range tests {
		if actual := getDepth(tt.param); actual != tt.res {
			t.Errorf(
				"got %d; expected %d,%+v",
				 actual, tt.res,tt.param	)
		}else{
			//t.Log("this ok,",tt.param)
			t.Logf("this ok,%+v",tt.param)
		}
	}
}


func TestTravers(t *testing.T){

	//fmt.Println(levelOrder(buildFullTree()))

	var a lgd
	fmt.Println(a.TraverseByLevel102(buildFullTree()))


	fmt.Println(a.reverseTraverse107(buildFullTree()))
}


func TestConverseToTree(t *testing.T){

	var e lgd
	root := e.converseSortedArrToBst([]int{-10, -3, 0, 5, 9  ,  17  , 117})
	fmt.Println(root)
	arr := []int{}
	e.TreeToArr(root,&arr)
	fmt.Println(arr)
}

//    1
// 2    5
//
func TestAct(t *testing.T){
	var e lgd
	arr := []int{}
	e.TreeToArr(buildBinSearTree(),&arr)

	fmt.Println(arr)
}


//               5
//       -3             17
// -10       0       9        117