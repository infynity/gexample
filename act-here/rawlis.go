package main

import (
	"container/list"
	"fmt"
)

func printList(coll *list.List) {
	for e := coll.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}

func printListR(coll *list.List) {
	for e := coll.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}

func main() {
	coll := list.New()

	coll.PushBack(1)
	coll.PushBack("Gopher")

	coll.PushFront("Cynhard")
	two := coll.PushFront(2)
	printList(coll)

	before2 := coll.InsertBefore("Before2", two)
	printList(coll)

	after2 := coll.InsertAfter("After2", two)
	printList(coll)

	fmt.Println("---------------")
	coll.MoveAfter(before2, two)
	printList(coll)

	coll.MoveBefore(after2, two)
	printList(coll)

	coll.MoveToFront(before2)
	printList(coll)
	coll.MoveToBack(after2)
	printList(coll)

	fmt.Println("===================")
	coll2 := list.New()
	coll2.PushBack(3)
	coll2.PushFront("Tomcat")

	coll.PushBackList(coll2)
	coll.PushFrontList(coll2)

	printList(coll)

	fmt.Println("~~~~~~~~~~~~~~~~~~")
	fmt.Println(coll.Front().Value,"----")
	fmt.Println(coll.Back().Value,"====")

	fmt.Println(coll.Len(),"+++++")

	coll.Remove(two)

	printList(coll)

	coll.Init()
	fmt.Println(123)
	printList(coll)
}


func v103stat(){
	model.GetSlaveDB().
		Model(model.Pc{}).
		Order("id asc").
		Offset(10*i).
		Limit(10).
		Scan(&pcs)
	var ids []int
	for _, pc := range pcs {
		fmt.Println(pc)
		ids = append(ids, pc.ID)
	}
	areaMap, _ := model.GetPcsAreaMap(ids)

	//fmt.Println(pcs)
	fmt.Printf("%+v",areaMap)

	for _, pc := range pcs {
		if val,ok:=areaMap[pc.ID];ok{
			fmt.Printf("%+v",val	)

			//存在区域的pc进行时长统计
			e.GetOnePcOnlineDuration(pc.ID,val.ID)

		}

		fmt.Println()

	}
}