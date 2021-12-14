package global

const (
    SUCCESS         = 200
    ERROR           = 500
    BAD_REQUEST     = 400
    UNAUTHORIZED    = 401

    ERROR_BIND_JSON = 10001
    ERROR_MYSQL     = 10002
    ERROR_REDIS     = 10003

    ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
    ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002


)

var MsgFlags = map[int]string{
    SUCCESS:                        "成功",
    ERROR:                          "失败",
    UNAUTHORIZED:                   "未用户认证",
    ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
    ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
    ERROR_BIND_JSON:                "解析请求体失败",
    ERROR_MYSQL:                    "操作MySQL数据库失败",
    ERROR_REDIS:                    "操作Redis数据库失败",
}

func GetMsg(code int) string {
    msg, ok := MsgFlags[code]
    if ok {
        return msg
    }

    return MsgFlags[ERROR]
}
