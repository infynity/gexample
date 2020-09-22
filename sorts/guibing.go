package sorts

import "fmt"

func merge(chan1,chan2 <-chan int)<-chan int{

	out :=make(chan int)


	go func() {
		v1,ok1  := <-chan1
		v2,ok2  := <-chan2

		for ok1||ok2{
			if !ok2 || (ok1 && v1<=v2){
				out<-v1
				v1,ok1=<-chan1

			}else{
				out<-v2
				v2,ok2=<-chan2

			}
		}

		close(out)

	}()

	return out
}


func MergeN(inputs ...<-chan int)<-chan int{

	if len(inputs) == 1 {
		return inputs[0]
	}


	m:= len(inputs)/2

	return merge(MergeN(inputs[:m]...),MergeN(inputs[m:]...))
}

func Mergetest() {
	ch := pipeline.Merge(
		pipeline.InMemSort(
			pipeline.ArraySource(100, 13, 15, 17, 123, 321)),
		pipeline.InMemSort(
			pipeline.ArraySource(2, 4, 6, 8)))
	fmt.Println(ch)
	for v := range ch {
		fmt.Println("paixuhou-------", v)
	}
}