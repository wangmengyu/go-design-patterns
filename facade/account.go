package facade

import "fmt"

type Account struct {
	Name string //账号主人
}

func NewAccount(name string) *Account {
	return &Account{
		Name: name,
	}
}

//检查账户是不是正确
func (a *Account) CheckAccount(accountName string) error {
	if a.Name != accountName {
		return fmt.Errorf("Account Name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}
