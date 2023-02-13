package main

import (
	"design_patterns/struct/facade/wallet"
	"fmt"
	"log"
)

func main() {
	fmt.Println()
	/**
	外观是一种结构型设计模式， 能为复杂系统、 程序库或框架提供一个简单 （但有限） 的接口。
	尽管外观模式降低了程序的整体复杂度， 但它同时也有助于将不需要的依赖移动到同一个位置。
	 */
	walletFacade := wallet.NewWalletFacade("abc", 1234)
	fmt.Println()

	err := walletFacade.AddMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.DeductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
