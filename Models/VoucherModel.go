package Models

type Voucher struct {
	Id      		uint   `json:"id"`
	SerialNumber    string `json:"serialNumber"`
	VoucherValidity string `json:"voucherValidity"`
	ServiceId   	string `json:"serviceId"`
	PackageName 	string `json:"packageName"`
	Status 			string `json:"status"`
}
func (b *Voucher) TableName() string {
	return "voucher"
}

/*func (u *Voucher) AfterFind(tx *gorm.DB) (err error) {
	u.SerialNumber = Service.Decrypt(u.SerialNumber)
	return
}*/
