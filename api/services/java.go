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
	javaURL = "http://10.11.65.65:8080/receive_json"
)

func SendCodeToJava(code Task, context *gin.Context) {
	jsonBytes, err := json.Marshal(code)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "bad code"})

	}

	resp, err := http.Post(javaURL, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
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
	defer resp.Body.Close()
	responseBody := string(bodyBytes)
	fmt.Println(responseBody)
	// Return the response to the client
	context.JSON(http.StatusOK, gin.H{
		"response": responseBody,
	})
	defer resp.Body.Close()

}
