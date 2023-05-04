package e

const (
	SUCCESS       = 200
	InvalidParams = 400
	AccessDenied  = 403
	ERROR         = 500

	ErrorUserExist             = 1001
	ErrorUserNotExist          = 1002
	ErrorWrongPassword         = 1003
	ErrorPasswordEncrypt       = 1004
	ErrorInvalidPasswordLength = 1005

	ErrorTokenGenerate = 2001
	ErrorTokenMissing  = 2002
	ErrorTokenParsing  = 2003
	ErrorTokenExpired  = 2004

	ErrorOpenImageFile   = 3001
	ErrorSaveImageFile   = 3002
	ErrorDecodeImageFile = 3003

	ErrorQuestionNotExist = 4001

	ErrorDataBase   = 5001
	ErrorCache      = 5002
	FileSystemError = 5003

	ErrorAnswerNotExist = 6001
	ErrorAlreadyLiked   = 6002
)

var MsgFlags = map[int]string{
	SUCCESS:                    "ok",
	ERROR:                      "fail",
	InvalidParams:              "请求参数错误",
	ErrorDataBase:              "数据库错误",
	ErrorUserExist:             "用户已经存在",
	ErrorUserNotExist:          "用户不存在",
	ErrorWrongPassword:         "用户密码错误",
	ErrorPasswordEncrypt:       "密码加密错误",
	ErrorInvalidPasswordLength: "密码长度错误",
	ErrorTokenGenerate:         "token签发错误",
	ErrorTokenMissing:          "验证身份时token缺失",
	ErrorTokenParsing:          "验证身份时token解析失败",
	ErrorTokenExpired:          "token已过期",
	ErrorOpenImageFile:         "无法打开图像文件",
	ErrorSaveImageFile:         "无法保存图像文件",
	FileSystemError:            "服务器文件系统错误",
	ErrorDecodeImageFile:       "服务器无法解析图像文件",
	ErrorQuestionNotExist:      "问题不存在",
	ErrorAnswerNotExist:        "回答不存在",
	AccessDenied:               "无权限访问",
	ErrorCache:                 "缓存错误",
	ErrorAlreadyLiked:          "无法重复点赞",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
