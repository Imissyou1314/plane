package models

import (
	"time"
)

// Result 返回模块
type Result struct {
	ResultInfo   string
	ResultCode   int
	ResultStatus bool
	ResultData   map[string]interface{}
	ResultError  interface{}
	ResultDate   time.Time
}

// GetResultModel 获取返回模型
func GetResultModel() (result *Result) {
	result = &Result{}
	result.ResultDate = time.Now().UTC()
	return
}

// SetResultModel 设置返回值
func SetResultModel(key, info string, data interface{}, err error) (result *Result) {
	result = &Result{}
	result.ResultDate = time.Now().UTC()
	if err != nil {
		result.ResultError = err
		result.ResultCode = 500
		result.ResultStatus = false
	} else {
		result.ResultData = SetMapData(key, data)
		result.ResultStatus = true
		result.ResultCode = 200
	}
	return
}

// SetMapData 封装Map对象作为返回值
func SetMapData(key string, data interface{}) (mapData map[string]interface{}) {
	mapData = make(map[string]interface{})
	mapData[key] = data
	return
}

// GetMapData 获取Map进行填充数据
func GetMapData() (result map[string]interface{}) {
	result = make(map[string]interface{})
	return
}

// SetResultWithMapData 多值返回
func SetResultWithMapData(data map[string]interface{}, err error) (result *Result) {
	result = &Result{}
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
	return
}

// SetWithErr 错误返回
func SetWithErr(err error) (result *Result) {
	result = SetResultModel("", "发成错误！", nil, err)
	return
}
