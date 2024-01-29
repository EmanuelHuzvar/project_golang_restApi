package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

var (
	csharpURL = "http://10.11.65.65:8080/receive_json"
)

type Task struct {
	Language string `json:"language"`

	Code string `json:"code"`
}

func SendCodeToCSharp(code Task, context *gin.Context) {

	jsonBytes, err := json.Marshal(code)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "bad code"})

	}

	resp, err := http.Post(csharpURL, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		_ = context.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	fmt.Println(resp)
	bodyBytes, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read response body",
		})
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	responseBody := string(bodyBytes)
	fmt.Println(responseBody)
	// Return the response to the client
	context.JSON(http.StatusOK, gin.H{
		"response": responseBody,
	})
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

}
