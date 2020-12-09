package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"sync"
)
type NSQHandler struct {
}

func (this *NSQHandler) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}

func testNSQ() {
	waiter := sync.WaitGroup{}
	waiter.Add(1)

	go func() {
		defer waiter.Done()
		config:=nsq.NewConfig()
		config.MaxInFlight=9

		//建立多个连接
		for i := 0; i<3; i++ {
			consumer, err := nsq.NewConsumer("test30", "struggle-test-manyconsumer", config)
			if nil != err {
				fmt.Println("err", err)
				return
			}

			consumer.AddHandler(&NSQHandler{})
			err = consumer.ConnectToNSQD("10.20.22.83:6205")
			if nil != err {
				fmt.Println("err", err)
				return
			}
		}
		select{}

		fmt.Println("done")
	}()

	waiter.Wait()
}
func main() {
	testNSQ()

}