package errno

// 自定义的错误码
const (
	Success       = 200
	Error         = 300
	InvalidParams = 400
)

// 错误码对应的错误消息
var Msg = map[int]string{
	Success:       "成功",
	Error:         "错误",
	InvalidParams: "请求参数错误",
}
