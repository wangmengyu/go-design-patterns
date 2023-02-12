package main

import (
	"design_patterns/behavioral/template/otp"
)

func main() {
	smsOTP := &otp.Sms{}
	o := otp.Otp{
		IOtp: smsOTP,
	}
	err := o.GenAndSendOTP(4)
	if err != nil {
		return
	}

	emailOTP := &otp.Email{}
	o = otp.Otp{
		IOtp: emailOTP,
	}
	err = o.GenAndSendOTP(4)
	if err != nil {
		return
	}
}
