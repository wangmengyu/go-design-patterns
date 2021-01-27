package main

import (
	"fmt"
	"sync"
)

var once sync.Once // 通常是用这个的. sync.Once来创建单例
type Single struct {
}

var singleInstance *Single

func getInstance() *Single {

	if singleInstance == nil {
		once.Do(func() {
			fmt.Println("Creating single instance now.")
		})
	} else {
		fmt.Println("Single instance already created")
	}

	return singleInstance
}
func main() {

	for i := 0; i < 30; i++ {
		go getInstance()
	}

}
