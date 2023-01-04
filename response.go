package main

import "github.com/gin-gonic/gin"

var ErrorCode = map[int]string{
	0:   "",
	1201: "Params err",
}

type BaseResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ErrorResponse struct {
	*BaseResponse
	Data string `json:"data"`
}

func Response(c *gin.Context, data interface{}) {
	c.JSON(200, data)
}

func Error(c *gin.Context, code int) {
	c.JSON(200, ErrorResponse{
		BaseResponse: &BaseResponse{
			Code: code,
			Msg:  ErrorCode[code],
		},
		Data: "",
	})
}

func ErrorString(c *gin.Context, errorStr string) {
	c.JSON(200, ErrorResponse{
		BaseResponse: &BaseResponse{
			Code: 1226,
			Msg:  errorStr,
		},
		Data: "",
	})
}

func ErrorExcept(c *gin.Context, errorCode int,errorStr string) {
	c.JSON(200, ErrorResponse{
		BaseResponse: &BaseResponse{
			Code: errorCode,
			Msg:  errorStr,
		},
		Data: "",
	})
}