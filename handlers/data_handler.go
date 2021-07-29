package handlers

import (
	"csat-report-webhook/entities"
	"csat-report-webhook/viewmodels"
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
	"os"
	"strconv"

	"csat-report-webhook/utils/helper"

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

	reqBody, err := helper.MakeMetabaseRequest(jsonBody, *h.multichannel)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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

func (h *DataHandler) FormSubmissionData(c *gin.Context) {
	var jsonBody viewmodels.JSONRequest
	err := c.ShouldBindJSON(&jsonBody)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cardID := os.Getenv("FORM_SUBMISSION_CARD_ID")

	reqBody, err := helper.MakeMetabaseRequest(jsonBody, *h.multichannel)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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

func (h *DataHandler) FormSentWithFormat(c *gin.Context) {
	format := c.Param("format")

	var jsonBody viewmodels.JSONRequest
	err := c.ShouldBindJSON(&jsonBody)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cardID := os.Getenv("FORM_SENT_CARD_ID")

	metabase := entities.NewMetabase()

	intCardID, err := strconv.Atoi(cardID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	serializedParams, err := helper.MakeMetabaseSerializedParams(jsonBody, *h.multichannel)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	metabaseResponse, httpResp := metabase.Login().GetQuestionWithFormat(intCardID, serializedParams, format)

	_, params, err := mime.ParseMediaType(httpResp.Header.Get("Content-Disposition"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, params["filename"]))
	c.Data(http.StatusOK, httpResp.Header.Get("Content-Type"), metabaseResponse)
	return
}

func (h *DataHandler) FormSubmissionWithFormat(c *gin.Context) {
	format := c.Param("format")

	var jsonBody viewmodels.JSONRequest
	err := c.ShouldBindJSON(&jsonBody)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cardID := os.Getenv("FORM_SUBMISSION_CARD_ID")

	metabase := entities.NewMetabase()

	intCardID, err := strconv.Atoi(cardID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	serializedParams, err := helper.MakeMetabaseSerializedParams(jsonBody, *h.multichannel)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	metabaseResponse, httpResp := metabase.Login().GetQuestionWithFormat(intCardID, serializedParams, format)

	_, params, err := mime.ParseMediaType(httpResp.Header.Get("Content-Disposition"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, params["filename"]))
	c.Data(http.StatusOK, httpResp.Header.Get("Content-Type"), metabaseResponse)
	return
}
