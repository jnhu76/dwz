package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",

	ERROR_EXIST_URL:       "已存在该URL",
	ERROR_EXIST_URL_FAIL:  "获取已存在URL失败",
	ERROR_NOT_EXIST_TAG:   "该URL不存在",
	ERROR_GET_URL_FAIL:    "获取URL失败",
	ERROR_COUNT_URL_FAIL:  "统计URL失败",
	ERROR_ADD_URL_FAIL:    "新增URL失败",
	ERROR_EDIT_URL_FAIL:   "修改URL失败",
	ERROR_DELETE_URL_FAIL: "删除URL失败",
	ERROR_EXPORT_URL_FAIL: "导出URL失败",
	ERROR_IMPORT_URL_FAIL: "导入URL失败",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
}

// GetMsg get error informathion based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
