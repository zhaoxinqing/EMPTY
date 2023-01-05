package app

import (
	"fmt"
	"log"
	"time"
)

func A() {

	botSignChannel := make(chan int)
	accountChannel := make(chan int)

	for {
		select {
		case orderIndex := <-botSignChannel:
			// global.Log.Info("====> botSignChannel pop index_[%v]_orderid[%v]", orderIndex, globalHyOrderList[orderIndex].ID)
			if accountIndex, ok := <-accountChannel; ok {
				// global.Log.Info("====> accountChannel pop index_[%v]_address[%v]", accountIndex, robotAccountList[accountIndex].Addr)
				go func() {
					defer func() {
						if err := recover(); err != nil {
							log.Println(err)
						}
						time.Sleep(100 * time.Millisecond)
						accountChannel <- accountIndex
						// global.Log.Info("====> accountChannel push index_[%v]_address[%v]", accountIndex, robotAccountList[accountIndex].Addr)
					}()
					// make bot sign
					fmt.Println(orderIndex)
				}()
			}
		case <-time.After(time.Second * 12):
			b()
		}
	}
}

func b() {}

// go func() {}
// 开启一个子协程执行该匿名函数内的逻辑
// go func()+defer一般用在：多次远程调用，每一次新开一个协程去执行，无需等待结果即可循环下一次

// defer
// 捕获异常
// 封装一些必须要执行的逻辑，顺序执行情况下担心前面代码return了而忘记/走不到该行逻辑。比如锁的释放。

// select
// 用于多个channel的结合，这些channel会通过类似于are-you-ready polling的机制来工作。
// 一旦某一个channel准备好接/发数据，select就会完成该对应的代码块内容
