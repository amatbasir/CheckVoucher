package Models

import Config "CheckVoucherStatus/DBConfig"

func GetVoucherBySN(voucher *Voucher, sn string)(err error)  {
	if err = Config.DB.Where("serial_number=?", sn).First(voucher).Error; err != nil{
		return err
	}
	return nil
}