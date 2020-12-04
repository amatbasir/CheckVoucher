package Controller

import (
	"CheckVoucherStatus/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)


// Get Users godoc
// @Summary Show Voucher Detail
// @Description get voucher detail by Serial Number
// @Tags voucher
// @Accept  json
// @Produce  json
// @Param sn path string true "Serial Number"
// @Success 200 {object} Models.User
// @Router /voucher/info [get]
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
