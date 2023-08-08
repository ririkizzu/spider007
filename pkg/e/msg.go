package e

type Code struct {
	code int
}

var MsgFlags = map[int]string{
	SUCCESS:        "请求成功",
	INVALID_PARAMS: "请求参数错误",
	ERROR:          "请求错误",

	MYSQL_CONN_ERR: "MySQL连接失败",
	MYSQL_ERR:      "MySQL操作异常",
}

func GetMsg(code int) string {
	if msg, ok := MsgFlags[code]; ok {
		return msg
	}
	return MsgFlags[ERROR]
}
