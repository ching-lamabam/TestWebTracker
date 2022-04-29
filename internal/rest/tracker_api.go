package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ching-lamabam/web-tracker/internal/pkg"
)

type trackerRest struct {
	trackerSvc pkg.TrackerSvc
}

func New() trackerRest {
	api := trackerRest{}
	api.trackerSvc = pkg.NewSvc()
	return api
}

func (rest *trackerRest) GetPing(c *gin.Context) {
	status := rest.trackerSvc.GetPing()
	if status == 1 {
		c.JSON(http.StatusOK, "OK")
	} else {
		c.Status(http.StatusServiceUnavailable)
	}
}

func (rest *trackerRest) GetImg(c *gin.Context) {
	path := rest.trackerSvc.GetImg()
	if len(path) > 0 {
		c.File(path)
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusServiceUnavailable)
	}
}
