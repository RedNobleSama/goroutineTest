/**
* @Author: oreki
* @Date: 2021/6/5 11:08
* @Email: a912550157@gmail.com
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

var t = make(chan bool)
var b = make(chan bool)

func TimeSub() int {
	Timestart := time.Now()
	//Timeend,_ := time.Parse("2006-01-02 15:25:05", "2021-09-08 15:05:00")
	targetTime := "2021-09-08 17:05:00" //待转化为时间戳的字符串
	timeLayout := "2006-01-02 15:04:05" //转化所需模板，go默认时间
	//go默认时间 很好记  6 1 2 3 4 5
	loc, _ := time.LoadLocation("Local") //获取本地时区
	Timeend, _ := time.ParseInLocation(timeLayout, targetTime, loc)

	left := int(Timeend.Sub(Timestart).Seconds())
	fmt.Println(left)
	return left
}

func uptime(wg *sync.WaitGroup) {
	time.Sleep(time.Second * 20)
	t <- true
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	ch := make(chan bool)
	var a bool
	delayTime := TimeSub()
	go Time(time.Duration(delayTime), ch, &wg)
	go uptime(&wg)
	go func() {
		if <-ch {
			a = true
		} else {
			a = false
		}
		fmt.Println(a)
		wg.Done()
	}()
	wg.Wait()
}

func Time(delayTime time.Duration, ch chan bool, wg *sync.WaitGroup)  {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	timeAfterTrigger := time.After(time.Second * delayTime)
	select {
		case curTime := <-timeAfterTrigger:
			fmt.Println(curTime.Format("2006-01-02 15:04:05"))
			ch <- true
			defer wg.Done()
		case <- t:
			//runtime.Goexit()
			fmt.Printf("中止时间 %s\n", time.Now().Format("2006-01-02 15:04:05"))
			go Time(delayTime, ch, wg)
	}

}
