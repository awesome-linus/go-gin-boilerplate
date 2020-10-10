package infrastructure

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateJsonResponse(c *gin.Context, statusCode int, payload interface{}) {
	// res, err := json.MarshalIndent(payload, "", "    ")
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, []byte(err.Error()))
	// 	return
	// }
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(statusCode, payload)
}

// respondError レスポンスとして返すエラーを生成する
func CreateErrorResponse(c *gin.Context, err error) {
	// logger := CreateLogger()
	// logger.Error(err.Error(), zap.String("RequestID", middleware.GetReqID(r.Context())))

	hc := &HTTPErrorCreator{}
	he := hc.CreateFromMsg(err.Error())
	CreateJsonResponse(c, he.Code, he)
}

// HTTPError エラー用
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (he *HTTPError) Error() string {
	return fmt.Sprintf("code=%d, message=%v", he.Code, he.Message)
}
