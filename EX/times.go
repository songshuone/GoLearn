package main

import (
	"time"
	"fmt"
	"math/rand"
)
//定时器
func main() {
	//timer1:=time.NewTimer(time.Second*2)
	//
	//fmt.Println(<-timer1.C)
	//
	//fmt.Println("Timer 1 expired")
	//
	//timer2:=time.NewTimer(time.Second)
	//go func() {
	//	<-timer2.C
	//	fmt.Println("Timer 2 expired")
	//}()
	//stop2:=timer2.Stop()
	//if stop2 {
	//	fmt.Println("Timer 2 stoped")
	//}
	fmt.Println(time.Now().Nanosecond())
	r:=	rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(r.Intn(1000))
}

