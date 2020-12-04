package Routes

import (
	"CheckVoucherStatus/Controller"
	"github.com/gin-gonic/gin"
)
func SetupRouter() *gin.Engine {
	r := gin.Default()

	usergroup := r.Group("/voucher")
	{
		usergroup.GET("info", Controller.GetVoucherBySN)

	}
	return r
}