package main

import "fmt"

type St struct {
	Top int

	Bot  int
	Max  int
	Body []int
}

func (st *St) push(val int) (bl bool) {
	if st.Top == st.Max+1 {
		bl = false
		return
	}
	st.Body[st.Top] = val
	if st.Top != st.Max {
		st.Top++
	}
	bl = true
	return
}

func (st *St) pull() (re int, bl bool) {
	if st.Top == st.Bot {
		bl = false
		return
	}
	re = st.Body[st.Top]
	st.Top--
	bl = true
	return
}


func main() {
	t := []int{0, 1, 2, 3, 5, 6}
	st := St{
		Top:  5,
		Body: t,
		Bot:  0,
		Max:  5,
	}
	st1 := St{
		Top:  1,
		Body: t,
		Bot:  0,
		Max:  5,
	}

	a := []int{0, 7, 8, 9, 10, 11}
	st.Body = a
	pullPro(&st, &st1)
}

func pullPro(st *St, st1 *St) (re int) {
	jp := true
	var reo int
	for jp {
		reo, jp = st.pull()
		if jp {
			st1.push(reo)
		}
	}

	jp = true
	for jp {
		reo, jp = st1.pull()
		if jp {
			fmt.Println(reo)
		}
	}
	return
}
