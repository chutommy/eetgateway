package server

import (
	"errors"
	"net/http"

	"github.com/chutommy/eetgateway/pkg/eet"
	"github.com/gin-gonic/gin"
)

// ErrUnexpectedFailure is returned if no error was expected but some did occur.
var ErrUnexpectedFailure = errors.New("unexpected error")

// Handler provides handling options for incoming requests.
type Handler interface {
	HTTPHandler() http.Handler
}

type handler struct {
	gatewaySvc eet.GatewayService
}

func (h *handler) HTTPHandler() http.Handler {
	return h.ginEngine()
}

// NewHandler returns an HTTP Handler implementation.
func NewHandler(gatewaySvc eet.GatewayService) Handler {
	return &handler{
		gatewaySvc: gatewaySvc,
	}
}

func (h *handler) ginEngine() *gin.Engine {
	r := gin.New()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", h.ping)
		v1.POST("/eet/:certID", h.eet)
	}

	return r
}

func (h *handler) ping(c *gin.Context) {
	if err := h.gatewaySvc.Ping(); err != nil {
		if errors.Is(err, eet.ErrMFCRConnection) {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": eet.ErrMFCRConnection.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrUnexpectedFailure.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusText(http.StatusOK)})
}

func (h *handler) eet(c *gin.Context) {
	certID := c.Param("certID")

	var req *HTTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	odpoved, err := h.gatewaySvc.Send(c, certID, encodeRequest(req))
	if err != nil {
		if errors.Is(err, eet.ErrCertificateRetrieval) {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": eet.ErrCertificateRetrieval.Error()})
			return
		} else if errors.Is(err, eet.ErrRequestConstruction) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": eet.ErrRequestConstruction.Error()})
			return
		} else if errors.Is(err, eet.ErrMFCRConnection) {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": eet.ErrMFCRConnection.Error()})
			return
		} else if errors.Is(err, eet.ErrMFCRResponseParse) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": eet.ErrMFCRResponseParse.Error()})
			return
		} else if errors.Is(err, eet.ErrMFCRResponseVerification) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": eet.ErrMFCRResponseVerification.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrUnexpectedFailure.Error()})
		return
	}

	c.JSON(http.StatusOK, decodeResponse(odpoved))
}