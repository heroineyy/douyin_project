package middleware

import (
	"byte_douyin_project/common"
	"crypto/sha1"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"regexp"
	"unicode"
)

func SHA1(s string) string {

	o := sha1.New()

	o.Write([]byte(s))

	return hex.EncodeToString(o.Sum(nil))
}

// 验证密码强度（至少6位，包含字母和数字）
func validatePassword(password string) bool {
	if len(password) < 6 {
		return false
	}

	hasLetter := false
	hasDigit := false

	for _, char := range password {
		if unicode.IsLetter(char) {
			hasLetter = true
		} else if unicode.IsDigit(char) {
			hasDigit = true
		}
	}

	return hasLetter && hasDigit
}

// 验证用户名格式（只允许字母、数字、下划线）
func validateUsername(username string) bool {
	if len(username) < 3 || len(username) > 20 {
		return false
	}
	validUsername := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	return validUsername.MatchString(username)
}

func SHAMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		username := context.Query("username")
		// 验证用户名格式
		if !validateUsername(username) {
			common.ErrorResponse(context, "用户名格式错误：只能包含字母、数字和下划线，长度3-20位")
			return
		}
		password := context.PostForm("password")
		if !validatePassword(password) {
			common.ErrorResponse(context, "密码强度不足：至少6位，需包含字母和数字")
			context.Abort() //阻止执行
			return
		}
		context.Set("password", SHA1(password))
		context.Next()
	}
}
