package models

import (
	"time"
)

/**
 * 返回类型
 */
type Result struct {
	ResultInfo   string
	ResultCode   int
	ResultStatus bool
	ResultData   interface{}
	ResultError  interface{}
	ResultDate   time.Time
}

func GetResultModel() *Result {
	result := &Result{}
	result.ResultDate = time.Now().UTC()
	return result
}

func SetResultModel(data interface{}, err error) *Result {
	result := &Result{}
	result.ResultDate = time.Now().UTC()
	if err != nil {
		result.ResultError = err
		result.ResultCode = 500
		result.ResultStatus = false
	} else {
		result.ResultData = data
		result.ResultStatus = true
		result.ResultCode = 200
	}
	return result
}

func SetWithErr(err error) *Result {
	result := SetResultModel(nil, err)
	return result
}
