package middlewares

import (
	"bytes"
	"csat-report-webhook/entities"
	"csat-report-webhook/handlers"
	"csat-report-webhook/viewmodels"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MultichannelAuth(dataHandler *handlers.DataHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		adminMulchanEmail := c.GetHeader("Authorization-Email")
		adminMulchanPassword := c.GetHeader("Authorization-Password")

		requestBody, _ := json.Marshal(map[string]string{
			"email":    adminMulchanEmail,
			"password": adminMulchanPassword,
		})

		resp, err := http.Post("https://multichannel.qiscus.com/api/v1/auth", "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			log.Fatalf("An Error Occured: %v", err)
			return
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			unauthorizedRespBody := map[string]string{
				"message": "Unauthorized",
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, unauthorizedRespBody)
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("An Error Occured: %v", err)
			return
		}

		var multichannelAuthResp viewmodels.MultichannelAuthResponse
		json.Unmarshal(body, &multichannelAuthResp)

		multichannel := entities.NewMultichannel(multichannelAuthResp.Data.User.App.AppCode, multichannelAuthResp.Data.User.App.SecretKey)
		dataHandler.SetMultichannel(multichannel)
	}
}
