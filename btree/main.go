package main

import (
	"log"
	"ruok/btree/Tree"
	"fmt"
)

func main()  {

    root:=Tree.NewBTree(13)
//         13
//    11         12
//  9   10     7    8
//               6

	root.ConnectLeft(11).ConnectRight(12).String()
    {
    	root.Left.ConnectLeft(9).ConnectRight(10)
		root.Right.ConnectLeft(7).ConnectRight(8)
    	root.Right.Right.ConnectLeft(6)
		root.Right.Right.Left.ConnectLeft(5)
	}


	log.Println(struct {

	}{})
    fmt.Println(root.Level())
    root.Preorder()
    fmt.Println()
    root.Midorder()










}


