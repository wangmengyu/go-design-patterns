package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{} // 锁
type Single struct {
}

var singleInstance *Single // 全局变量
func getInstance() *Single { // 这个方法不是golang 常用的, 但是可以工作.  check-lock-check 方法
	if singleInstance == nil {
		lock.Lock()         // 开始加锁
		defer lock.Unlock() // 方法结束之前解锁
		if singleInstance == nil {
			//这里判断是空的. 可以进行创建.
			fmt.Println("Creating single instance now...")
			singleInstance = &Single{} //

		} else {
			fmt.Println("Single instance already created.")
		}
	}
	return singleInstance
}

func main() {
	//并发30次获得当前实例
	for i := 0; i < 30; i++ {

		go getInstance()
	}
	_, _ = fmt.Scanln()
}
