package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	t := time.Now()
	p(t)
	p(t.Format(time.RFC3339))

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")


	timestamp:=t1.Unix()

	p(timestamp)
	p(t1,999,t1.Unix())
	timeLayout := "2006-01-02 15:04:05"  //转化所需模板

	//location, _ := time.LoadLocation("Asia/Shanghai")
	var cstZone = time.FixedZone("CST", 0)
	datetime := time.Unix(timestamp, 0).In(cstZone).Format(timeLayout)

	p(datetime,101010)

	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	//t2, e := time.Parse(form, "8 42 PM")
	t2, e := time.Parse(form, "8 42 PM")
	p(t2)
	p()

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}