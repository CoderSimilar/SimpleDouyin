package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LarryTestResponse struct {
	Message  string      `json:"message"`
	Response interface{} `json:"response"`
	Err      error       `json:"error"`
}

func LarryTest(c *gin.Context) {
	var response []User
	err := GetRowsByTable("users", &response)
	if err != nil {
		c.JSON(http.StatusOK, LarryTestResponse{
			Message:  "error",
			Response: response,
		})
		return
	} else {
		c.JSON(http.StatusOK, LarryTestResponse{
			Message:  "ok",
			Response: response,
		})
		return
	}

}
