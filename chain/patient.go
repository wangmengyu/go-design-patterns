package chain

/**
患者数据
*/
type Patient struct {
	Name              string
	RegistrationDone  bool // 是否注册
	DoctorCheckUpDone bool // 医生检查完毕
	MedicineDone      bool //配药完毕
	PaymentDone       bool // 支付完毕
}
