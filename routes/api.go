package route

import (
	"csat-report-webhook/entities"
	"csat-report-webhook/handlers"
	"csat-report-webhook/middlewares"

	"github.com/gin-gonic/gin"
)

type api struct {
	gin *gin.Engine
}

func NewApi(gin *gin.Engine) *api {
	return &api{gin}
}

func (a *api) Init() {
	var multichannel entities.Multichannel
	dataHandler := handlers.NewDataHandler(&multichannel)

	api := a.gin.Group("/api")
	{
		v1 := api.Group("/v1")

		v1.Use(middlewares.MultichannelAuth(dataHandler))
		{
			v1.POST("/form_sent", dataHandler.FormSentData)
			v1.POST("/form_submission", dataHandler.FormSubmissionData)
			v1.POST("/form_sent/:format", dataHandler.FormSentWithFormat)
			v1.POST("/form_submission/:format", dataHandler.FormSubmissionWithFormat)
		}
	}
}
