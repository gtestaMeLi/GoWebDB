package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type response struct {
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Response(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

func Success(c *gin.Context, status int, data interface{}) {
	Response(c, status, response{Data: data})
}

// NewErrorf creates a new error with the given status code and the message
// formatted according to args and format.
func Error(c *gin.Context, status int, format string, args ...interface{}) {
	err := errorResponse{
		Code:    strings.ReplaceAll(strings.ToLower(http.StatusText(status)), " ", "_"),
		Message: fmt.Sprintf(format, args...),
		Status:  status,
	}

	Response(c, status, err)
}

// func CreateRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
// 	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("token", "123456")

// 	return req, httptest.NewRecorder()
// }

// var (
// 	ErrInvId = errors.New("Error: invalid ID")
// )
