package tools

type replayMsg struct {
	Code   int    `json:"code"`
	SubMsg string `json:"sub_msg"`
}

type replayData struct {
	Data string `json:"data"`
}

//默认操作成功
func OperationSuccess() replayMsg {
	return replayMsg{Code: 200, SubMsg: "操作成功"}
}

//自定义msg操作成功
func OperationSuccessMsg(msg string) replayMsg {
	return replayMsg{Code: 200, SubMsg: msg}
}

//默认操作失败
func OperationFalse() replayMsg {
	return replayMsg{Code: 201, SubMsg: "操作失败"}
}

//自定义失败原因
func OperationFalseMsg(msg string) replayMsg {
	return replayMsg{Code: 201, SubMsg: "操作失败:" + msg}
}

//重复添加
func RepetitionFalse() replayMsg {
	return replayMsg{Code: 202, SubMsg: "重复添加"}
}

//返回数据为空
func ReturnDataNull() replayMsg {
	return replayMsg{Code: 203, SubMsg: "返回数据为空"}
}

//请登录
func PleaseLogin() replayMsg {
	return replayMsg{Code: 301, SubMsg: "请登录"}
}

//参数为空
func ParamNull() replayMsg {
	return replayMsg{Code: 401, SubMsg: "参数为空"}
}

//参数为空
func ParamNullMsg(msg string) replayMsg {
	return replayMsg{Code: 401, SubMsg: msg}
}

//默认参数错误
func ParamError() replayMsg {
	return replayMsg{Code: 402, SubMsg: "参数有误:"}
}

//自定义参数错误
func ParamErrorMsg(msg string) replayMsg {
	return replayMsg{Code: 402, SubMsg: "参数有误:" + msg}
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
