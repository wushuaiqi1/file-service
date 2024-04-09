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
