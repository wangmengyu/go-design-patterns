package facade

import "fmt"

type WalletFacade struct { // 钱包装饰器
	Account      *Account
	Wallet       *Wallet
	SecurityCode *SecurityCode // 安全码
	Notification *Notification // 通知
	Ledger       *Ledger       //账本
}

func NewWalletFacade(accountId string, code int) *WalletFacade {
	//开启一个新存着账号
	fmt.Println("Starting create account")
	wf := &WalletFacade{
		Account:      NewAccount(accountId),
		Wallet:       NewWallet(),
		SecurityCode: NewSecurityCode(code),
		Notification: &Notification{},
		Ledger:       &Ledger{},
	}

	return wf
}

func (w *WalletFacade) AddMoney(accountId string, securityCode int, amount int) error {

	fmt.Println("add money to wallet")
	err := w.Account.CheckAccount(accountId) //检测账号
	if err != nil {
		//检查账号是否有异常
		return err
	}
	//检查密码是否正确
	err = w.SecurityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}

	//增加金额
	w.Wallet.CreditBalance(amount)
	w.Notification.SendWalletCreditNotification()   // 发送金额变更通知
	w.Ledger.MakeEntry(accountId, "credit", amount) // 记录到账簿

	return nil

}

func (w *WalletFacade) DeductMoney(accountId string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet") // 从钱包扣除掉金额
	//检查账号是否正确

	err := w.Account.CheckAccount(accountId) //检测账号
	if err != nil {
		//检查账号是否有异常
		return err
	}
	//检查密码是否正确
	err = w.SecurityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}

	//从钱包扣除金额
	err = w.Wallet.DebitBalance(amount)
	if err != nil {
		return err
	}

	//发送通知
	w.Notification.SendWalletDebitNotification()
	//记录到账本
	w.Ledger.MakeEntry(accountId, "credit", amount)
	return nil

}
