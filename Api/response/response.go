package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseSuccess struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ResponseSuccess{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, ResponseError{
		Code:    code,
		Message: message,
	})
}

func SuccessWithMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, ResponseSuccess{
		Code:    200,
		Message: message,
		Data:    data,
	})
}

func ErrorWithMessage(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, ResponseError{
		Code:    code,
		Message: message,
	})
}
