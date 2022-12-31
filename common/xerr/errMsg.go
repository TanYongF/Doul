package xerr

var message map[uint32]string

func init() {
	//100xxx
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[SERVER_COMMON_ERROR] = "服务器开小差啦,稍后再来试一试"
	message[REUQEST_PARAM_ERROR] = "参数错误"
	message[TOKEN_EXPIRE_ERROR] = "token失效，请重新登陆"
	message[TOKEN_GENERATE_ERROR] = "生成token失败"
	message[DB_ERROR] = "数据库繁忙,请稍后再试"
	message[CACHE_ERROR] = "缓存繁忙,请稍后再试"
	message[DB_UPDATE_AFFECTED_ZERO_ERROR] = "更新数据影响行数为0"

	//200xxx
	message[NO_SUCH_USER] = "无此用户，请检查后重试！"
	message[WRONG_PASSWORD] = "登陆用户密码错误！请检查后重试！"
	message[NO_AUTH] = "用户未验证，请验证后再试"
	message[USERNAME_HAS_REGISTER] = "该用户名已被注册！"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
