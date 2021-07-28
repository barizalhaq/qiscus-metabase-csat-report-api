package handlers

import (
	"csat-report-webhook/entities"
	"csat-report-webhook/viewmodels"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type DataHandler struct {
	multichannel *entities.Multichannel
}

func NewDataHandler(multichannel *entities.Multichannel) *DataHandler {
	return &DataHandler{multichannel}
}

func (h *DataHandler) SetMultichannel(multichannel *entities.Multichannel) {
	h.multichannel = multichannel
}

func (h *DataHandler) FormSentData(c *gin.Context) {
	var jsonBody viewmodels.JSONRequest
	err := c.ShouldBindJSON(&jsonBody)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cardID := os.Getenv("FORM_SENT_CARD_ID")

	parameters := []viewmodels.MetabaseParameter{
		viewmodels.MetabaseParameter{
			Type: "category",
			Target: []interface{}{
				"variable",
				[]string{"template-tag", "app_code"},
			},
			Value: h.multichannel.GetAppID(),
		},
		viewmodels.MetabaseParameter{
			Type: "category",
			Target: []interface{}{
				"variable",
				[]string{"template-tag", "row_limit"},
			},
			Value: strconv.Itoa(jsonBody.Limit),
		},
	}
	if len(jsonBody.On) > 0 {
		target := []interface{}{
			"dimension",
			[]string{"template-tag", "created_at"},
		}
		parameters = append(parameters, viewmodels.MetabaseParameter{
			Type:   "date/all-options",
			Target: target,
			Value:  jsonBody.On,
		})
	}

	if len(jsonBody.StartDate) > 0 && len(jsonBody.EndDate) > 0 {
		layout := "2006-01-02"
		parsedStartDate, err := time.Parse(layout, jsonBody.StartDate)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		parsedEndDate, err := time.Parse(layout, jsonBody.EndDate)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		target := []interface{}{
			"dimension",
			[]string{"template-tag", "created_at"},
		}
		if !parsedEndDate.After(parsedStartDate) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "end_date must greater than start_date"})
			return
		}
		parameters = append(parameters, viewmodels.MetabaseParameter{
			Type:   "date/all-options",
			Target: target,
			Value:  fmt.Sprintf("%s~%s", jsonBody.StartDate, jsonBody.EndDate),
		})
	}
	reqBody := viewmodels.MetabaseDataRequest{
		IgnoreCache: true,
		Parameters:  parameters,
	}

	metabase := entities.NewMetabase()

	intCardID, err := strconv.Atoi(cardID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	metabaseResponse := metabase.Login().GetQuestionData(intCardID, jsonReqBody)

	c.JSON(http.StatusOK, gin.H{"data": metabaseResponse})
	return
}
