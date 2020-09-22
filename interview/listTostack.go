package main

import "fmt"

type D struct {
	Font int
	Real int
	Body []int
	Max  int
	Size int
}

func (d *D) pull() (re int, error bool) {

	if d.Font == d.Real {
		re = 0
		error = false
		return
	}

	re = d.Body[d.Font];
	d.Font = (d.Font + 1) % d.Max;
	d.Size--
	error = true
	return
}

func (d *D) push(val int) (error bool) {
	f := (d.Real + 1) % d.Max
	if f == d.Font {
		error = false
		return
	}
	index := d.Real % d.Max
	d.Body[index] = val
	d.Real = f
	d.Size++
	error = true
	return
}

func AllPush(d1 *D, d2 *D) {
	judge := true
	var re int
	for d1.Size > 0 {
		for judge {

			if d1.Size > 1 {
				re, judge = d1.pull();
				if !judge {
					break
				}
				judge = d2.push(re)
			} else {
				break
			}
		}

		re, _ = d1.pull()
		fmt.Println(re)
		judge = true
		for judge {
			if d2.Size > 0 {

				re, judge = d2.pull()

				if !judge {
					break
				}
				judge = d1.push(re)
			} else {
				break
			}
		}
	}
}

func main() {
	d1 := D{
		Font: 0,
		Max:  5,
		Real: 5,
		Size: 5,
		Body: []int{1, 2, 3, 4, 5},
	}

	d2 := D{
		Font: 0,
		Max:  5,
		Real: 0,
		Size: 0,
		Body: []int{0, 0, 0, 0, 0},
	}

	AllPush(&d1, &d2)
}