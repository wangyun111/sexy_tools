package tools

type replayMsg struct {
	Code   int    `json:"code"`
	SubMsg string `json:"sub_msg"`
}

type replayData struct {
	Data string `json:"data"`
}

//操作成功
func OperationSuccess() replayMsg {
	return replayMsg{Code: 200, SubMsg: "操作成功"}
}

//操作失败
func OperationFalse() replayMsg {
	return replayMsg{Code: 201, SubMsg: "操作失败"}
}

//自定义失败原因
func OperationFalseMsg(errmsg string) replayMsg {
	return replayMsg{Code: 201, SubMsg: errmsg}
}

//重复添加
func RepetitionFalse() replayMsg {
	return replayMsg{Code: 202, SubMsg: "重复添加"}
}

//返回数据为空
func ReturnDataNull() replayMsg {
	return replayMsg{Code: 204, SubMsg: "返回数据为空"}
}

//登录成功
func LoginSuccess() replayMsg {
	return replayMsg{Code: 206, SubMsg: "登录成功"}
}

//登录失败
func LoginFailure() replayMsg {
	return replayMsg{Code: 207, SubMsg: "登录失败"}
}

//请登录
func PleaseLogin() replayMsg {
	return replayMsg{Code: 208, SubMsg: "请登录"}
}

//参数错误
func ParamError(msg string) replayMsg {
	return replayMsg{Code: 403, SubMsg: msg}
}

//参数为空
func ParamNull() replayMsg {
	return replayMsg{Code: 402, SubMsg: "参数为空"}
}

// 自定义返回对象
func GetReturnObject(str string) replayData {
	data := replayData{Data: str}
	return data
}

//代码问题
func RoutineError(msg string) interface{} {
	return replayMsg{Code: 501, SubMsg: msg}
}
