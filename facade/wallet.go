package facade

import "fmt"

type Wallet struct {
	Balance int // 钱包类
}

func NewWallet() *Wallet {
	return &Wallet{}
}

func (w *Wallet) CreditBalance(amount int) {
	//存储金额
	w.Balance += amount
	fmt.Println("Wallet balance added successfully")

}

func (w *Wallet) DebitBalance(amount int) error {
	//余额是否足够
	if w.Balance < amount {
		return fmt.Errorf("Balance is not sufficient")
	}
	//扣除金额
	w.Balance -= amount
	fmt.Println("Wallet balance is Sufficient")
	return nil
}
