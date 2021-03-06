package httphandler

import (
	"net/http"
	"time"

	"github.com/chutommy/eetgateway/pkg/eet"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) sendSale(c *gin.Context) {
	// default request
	dateTime := eet.DateTime(time.Now())
	dateTime.Normalize()
	req := &SendSaleReq{
		UUIDZpravy:   eet.UUIDType(uuid.New().String()),
		DatOdesl:     &dateTime,
		PrvniZaslani: true,
		Overeni:      false,
		Rezim:        0,
	}

	// bind to default
	if err := c.ShouldBindJSON(&req); err != nil {
		err = bindingErr(err)
		c.JSON(http.StatusBadRequest, GatewayErrResp{err.Error()})
		_ = c.Error(err)
		return
	}

	req.DatOdesl.Normalize()
	req.DatTrzby.Normalize()

	odpoved, err := h.gateway.SendSale(c, req.CertID, []byte(req.CertPassword), sendSaleRequest(req))
	if err != nil {
		code, resp := gatewayErrResp(err)
		c.JSON(code, resp)
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, sendSaleResponse(req, odpoved))
}
