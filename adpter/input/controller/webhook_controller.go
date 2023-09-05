package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juliofilizzola/bot_discord/adpter/input/convert"
	"github.com/juliofilizzola/bot_discord/application/domain"
	"github.com/juliofilizzola/bot_discord/application/port/input"
)

func NewWebhookControllerInterface(serviceInterface input.WebhookDomainService) WebhookControllerInterface {
	return &webhookControllerInterface{
		service: serviceInterface,
	}
}

type WebhookControllerInterface interface {
	CreatePR(ctx *gin.Context)
}

type webhookControllerInterface struct {
	service input.WebhookDomainService
}

func (wb *webhookControllerInterface) CreatePR(ctx *gin.Context) {
	var body domain.Github

	err := ctx.Bind(&body)
	// ! this is alert from validation err.
	if err != nil {
		fmt.Println("err", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": "err from convert body",
		})
	}

	webhookId := ctx.Param("id")

	webhookToken := ctx.Param("token")

	dataGithub := convert.ConvertDomainGithub(&body)

	result := wb.service.Send(&dataGithub, webhookId, webhookToken, body.Action)

	ctx.JSON(http.StatusOK, gin.H{
		result: result,
	})
}
