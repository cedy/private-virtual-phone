package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cedy/private-virtual-phone/pkg/provider"
	"github.com/cedy/private-virtual-phone/pkg/storage"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Route struct {
	Path       string
	HTTPMethod string
	Handler    func(c *gin.Context)
}

type API struct {
	Routes []Route
	DB     storage.Storage
}

func (a *API) index() func(*gin.Context) {
	return a.sendSMSGET
}

func (a *API) sendSMSGET(c *gin.Context) {
	durationInPast := -1 * 7 * 24 * time.Hour // 7 days
	statuses := []provider.Status{provider.Delivered}
	to := time.Now().UTC()
	from := to.Add(durationInPast)
	messages, err := a.DB.GetMany(from, to, statuses)
	if err != nil {
		err := fmt.Errorf("error during getting messages from storage: %w", err)
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
}

func (a *API) sendSMSPOST(*gin.Context) {
}

func (a *API) receiveSmsGET(*gin.Context) {
}
