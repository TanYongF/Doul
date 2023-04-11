package main

import (
	"fmt"
)

func main() {
	//c := cron.New()
	////https://eddycjy.gitbook.io/golang/di-3-ke-gin/cron
	////sec min hour day_of_month mon day_of_week
	//c.AddFunc("* *  * * *", func() {
	//	log.Println("Run models.CleanAllTag...")
	//})
	//
	//c.Start()
	cc := make(chan int)
	fmt.Println("start...")
	cc <- 1
	fmt.Println("start2...")
	select {
	case msg1, ok := <-cc:
		if !ok {
			fmt.Println("error in ch1")
			return
		}
		fmt.Println(msg1)
	default:
		fmt.Println("select end")
	}

	fmt.Println("stop....")
	//t1 := time.NewTimer(time.Second * 200)
	//for {
	//	select {
	//	case <-t1.C:
	//		t1.Reset(time.Second * 10)
	//	}
	//}
}
