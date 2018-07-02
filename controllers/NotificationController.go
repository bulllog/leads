package controllers

import (
	"context"
	"net/http"
	"hubble.lead/model"
	"hubble.lead/utils"
)

func RecieveWebhookNotification(ctx context.Context,
	entity *model.NotificationEntity) utils.ResponseObj {
	response := utils.ResponseObj{
		IsError: true,
	}

	status, err := model.Insert(ctx, entity)

	if status == false {
		response.Response = err.Error()
		response.Status = http.StatusExpectationFailed
		return response
	}

	response.IsError = false
	response.Status = http.StatusOK
	return response
}
