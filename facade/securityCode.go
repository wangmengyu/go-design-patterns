package facade

import "fmt"

type SecurityCode struct {
	Code int
}

//创建新的密码
func NewSecurityCode(code int) *SecurityCode {
	return &SecurityCode{Code: code}
}

//检查密码是否正确
func (s *SecurityCode) CheckCode(code int) error {
	if s.Code != code {
		return fmt.Errorf("Security code is incorrect")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}
