package common

//通用返回结构体

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func OfSuccess(data any) Result {
	return Result{200, "success", data}
}

func OfFail(code int, msg string) Result {
	return Result{code, msg, nil}
}
