package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"hubble.lead/controllers"
)

type Host struct {
	Host string
}

//Host Controller
func RegisterHostHandler(c *gin.Context) {
	hostName := c.Param("hostName")
	ctx := appengine.NewContext(c.Request)
	log.Infof(ctx, "POST - HOST")
	response := controllers.RegisterHost(ctx, hostName)

	if response.IsError == false {
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusPreconditionFailed, response)
	}
}

//Host Controller
func UpdateHostHandler(c *gin.Context) {
	ctx := appengine.NewContext(c.Request)
	log.Infof(ctx, "PUT - HOST")
	hostName := c.Param("hostName")
	statusIn := c.Query("status")

	if statusIn == "" {
		c.JSON(http.StatusBadRequest, "status param not found")
		return
	}

	log.Infof(ctx, " Update status for host status " + statusIn)

	status := true
	if statusIn == "false" {
		status = false
	}


	response := controllers.ChangeStatus(ctx, hostName, status)

	if response.IsError == false {
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusPreconditionFailed, response)
	}
}

