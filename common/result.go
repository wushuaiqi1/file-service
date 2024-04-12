package common

//通用返回结构体

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Fail struct {
	Code int
	Msg  string
}

func OfSuccess(data any) Result {
	return Result{200, "success", data}
}

func OfFail(fail Fail) Result {
	return Result{fail.Code, fail.Msg, nil}
}

var SystemFail = Fail{10001, "系统繁忙"}
var UploadedFail = Fail{10002, "已上传，请勿重复操作"}
var MissingParam = Fail{10003, "缺少必要参数"}
var BodySizeLimit = Fail{10004, "文件过大，上传文件请小于8MB"}
