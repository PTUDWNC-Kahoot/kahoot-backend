package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

const (
	KeyContextID = "tracing_id"
	MsgSuccess   = "success"

	StatusSuccess = 1
	StatusFailure = 0
)

type ginHandler func(c *gin.Context) *Response

func GinWrap(handler ginHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := handler(c)
		resp.TraceID = cast.ToString(c.Request.Context().Value(KeyContextID))
		statusCode := resp.Status

		if resp.Status == 0 || resp.Status == 1 {
			statusCode = http.StatusOK
		}

		c.JSON(statusCode, resp)
	}
}

func SuccessWithData(result interface{}) *Response {
	return &Response{
		Status: StatusSuccess,
		Data: &Data{
			Result: result,
		},
		Message: MsgSuccess,
	}
}

func SuccessWithCollection(results interface{}, pagination *Pagination) *Response {
	return &Response{
		Status: StatusSuccess,
		Data: &Collection{
			Results:    results,
			Pagination: pagination,
		},
		Message: MsgSuccess,
	}
}

func Success() *Response {
	return &Response{
		Status:  StatusSuccess,
		Message: MsgSuccess,
	}
}

func Failure(err error) *Response {
	return &Response{
		Status:  StatusFailure, // @todo
		Message: err.Error(),
	}
}
