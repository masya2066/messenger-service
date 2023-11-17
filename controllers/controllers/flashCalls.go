package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"pager-service/models"
)

func Call(c *gin.Context) {

	var callBody models.Call
	serviceUrl := os.Getenv("URL_CALL_SERVICE") + "/validation/request"

	_ = c.ShouldBindJSON(&callBody)

	if callBody.Number != "" || callBody.Type != "" || callBody.Platform == "" {

		if string(callBody.Number[0]) != "+" {
			c.JSON(400, "error: number is incorrect. The first character must be '+'")
			return
		}
	} else {
		c.JSON(400, "error: number, type, platform (ios, android, desktop, web) are necessary params")
		return
	}

	payload := map[string]interface{}{
		"number":   callBody.Number,
		"type":     callBody.Type,
		"platform": callBody.Platform,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", serviceUrl, bytes.NewBuffer(payloadBytes))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", os.Getenv("API_KEY"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	var res models.CallResponse

	_ = json.NewDecoder(response.Body).Decode(&res)

	if res.ID == "" || res.PinHash == "" {
		c.JSON(400, gin.H{
			"error":         true,
			"service_error": "Incorrect request or service is incorrect",
		})
		return
	}

	c.JSON(200, gin.H{
		"success":  true,
		"response": res,
	})
}

func Verify(c *gin.Context) {
	var body models.Verify

	_ = c.ShouldBindJSON(&body)

	if len(body.Pin) < 4 || body.Pin == "" {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Incorrect pin or id",
		})
		return
	}

	reqBody := map[string]interface{}{
		"id":  body.ID,
		"pin": body.Pin,
	}

	payload, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", os.Getenv("URL_CALL_SERVICE")+"/validation/verify", bytes.NewBuffer(payload))
	req.Header.Set("Authorization", os.Getenv("API_KEY"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(response.Body)

	bd, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var resVerify models.VerifyResponse

	if err := json.Unmarshal(bd, &resVerify); err != nil {
		panic(err)
	}

	if resVerify.Number == "" {
		var errorRes models.ErrorVerifyLimit

		if err = json.Unmarshal(bd, &errorRes); err != nil {
			panic(err)
		}

		if errorRes.Code != 0 {
			c.JSON(400, errorRes)
			return
		}
		panic("error: error communicate with service")
	}

	if resVerify.Validated != true {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Incorrect pin code",
		})
		return
	}

	c.JSON(200, resVerify)
}
