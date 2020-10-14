package autil

import "testing"

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