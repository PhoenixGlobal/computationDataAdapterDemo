package main

import (
	"computationDataAdapterDemo/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func GetMatrix(shape int) string {
	rand.Seed(time.Now().UnixNano())
	matrix:=make([]int,0)
	for i:=0;i<shape;i++{
		number:=rand.Intn(100)
		matrix=append(matrix,number)
	}
	var matrixStr = make([]string, len(matrix))
	for k, v := range matrix {
		matrixStr[k] = fmt.Sprintf("%d", v)
	}
	var result = "[" + strings.Join(matrixStr, ",") + "]"
	return result
}


type GetDataResponse struct {
	*BaseResponse
	Data string `json:"data"`
}

// @Summary get data
// @Description
// @Accept  json
// @Produce  json
// @Param number query int false "number，0 is false，1 is true"
// @Param matrix query int false "shape of matrix"
// @Success  200 {object} GetDataResponse "ok"
// @Router /api/data [get]
func  GetData(c *gin.Context){
	isNumber:=c.Query("number")
	matrix:=c.Query("matrix")
	isNumberInt,_:=strconv.Atoi(isNumber)
	matrixInt,_:=strconv.Atoi(matrix)

	utils.MainLogger.Info("GetData params",
		zap.Any("isNumber", isNumber),
		zap.Any("matrix", matrix),
	)

	if isNumberInt==0 && matrixInt==0{
		utils.MainLogger.Error("GetData Params err",
			zap.Any("isNumber", isNumber),
			zap.Any("matrix", matrix),
		)
		Error(c, 1201)
		return
	}

	rand.Seed(time.Now().UnixNano())

	if isNumberInt!=0{
		number:=rand.Intn(8999)+1000
		utils.MainLogger.Debug("generate rand number",
			zap.Any("number", number),
		)
		numberStr:=strconv.Itoa(number)
		Response(c, GetDataResponse{
			BaseResponse: &BaseResponse{},
			Data: numberStr,
		})
		return
	}

	dataStr := "["
	for i:=0;i<matrixInt;i++{
		matrixStr:=GetMatrix(matrixInt)
		if i==matrixInt-1{
			dataStr=dataStr+matrixStr+"]"
		}else {
			dataStr=dataStr+matrixStr+","
		}
	}
	utils.MainLogger.Debug("generate rand matrix",
		zap.Any("dataStr", dataStr),
	)
	Response(c, GetDataResponse{
		BaseResponse: &BaseResponse{},
		Data: dataStr,
	})
	return
}