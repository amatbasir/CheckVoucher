package Controller

import (
	"CheckVoucherStatus/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetVoucherBySN(c *gin.Context)  {
	var voucher Models.Voucher
	sn := c.Params.ByName("serialNumber")
	err := Models.GetVoucherBySN(&voucher, sn)
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, voucher)
	}
}
