package Routes

import (
	"CheckVoucherStatus/Controller"
	"github.com/gin-gonic/gin"
)
func SetupRouter() *gin.Engine {
	r := gin.Default()

	vouchergroup := r.Group("/voucher")
	{
		vouchergroup.GET("info/:serialNumber", Controller.GetVoucherBySN)

	}
	return r
}