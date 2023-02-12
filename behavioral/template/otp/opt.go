package otp

type IOtp interface {
	GenRandomOTP(int) string       // 生成随机一次性密码
	SaveOTPCache(string)           //  保存密码缓存
	GetMessage(string) string      // 获得消息
	SendNotification(string) error //  发送通知
}

type Otp struct { // 基类
	IOtp IOtp //  包含 一个 具体实现接口 的结构体.
}

func (o *Otp) GenAndSendOTP(otpLength int) error {
	// 组合实际持有接口结构体的各个方法进行工作
	otp := o.IOtp.GenRandomOTP(otpLength)
	o.IOtp.SaveOTPCache(otp)
	message := o.IOtp.GetMessage(otp)
	err := o.IOtp.SendNotification(message)
	if err != nil {
		return err
	}
	return nil
}
