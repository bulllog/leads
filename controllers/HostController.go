package controllers

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
	"net/http"
	"time"
	"hubble.lead/utils"
	"hubble.lead/model"
	"hubble.lead/googlequeue"
)

func RegisterHost(ctx context.Context, host string) utils.ResponseObj {
	response := utils.ResponseObj{
		IsError:true,
	}
	hostReset := false
	if model.DoesHostExist(ctx, host) {
		log.Infof(ctx, "Host already exist " + host)
		response.IsError = false
		response.Response = "Host Reset Successfully"
		hostReset = true
	} else {
		log.Infof(ctx, "Host not exist " + host)
		if _, err := model.CreateHost(ctx, host); err == nil {
			hostReset = true
			response.IsError = false
			response.Response = "Host Registered"
		}

	}

	if hostReset == true {
		googlequeue.ResetSubscriberForHost(ctx, host)
	}
	return response
}

func ChangeStatus(ctx context.Context, hostName string, status bool) utils.ResponseObj {
	response := utils.ResponseObj{
		IsError:true,
	}

	updatedHost := model.HostEntity{
		hostName,
		status,
		time.Now(),
	}

	if model.DoesHostExist(ctx, hostName) {
		log.Infof(ctx, "Host already exist " + hostName)
		if model.UpdateHost(ctx, &updatedHost) {
			response.IsError = false
			response.Response = "Host Reset Successfully"
		}
		return response
	} else {
		log.Infof(ctx, "Host not exist " + hostName)
		response.IsError = false
		response.Response = "Host does not found"
		response.Status = http.StatusNotFound
	}
	return response

}



