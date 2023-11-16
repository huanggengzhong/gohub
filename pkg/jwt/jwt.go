package jwt

import (
	"github.com/gin-gonic/gin"
	jwtpkg "github.com/golang-jwt/jwt"
	"gohub/pkg/app"
	"gohub/pkg/config"
	"gohub/pkg/logger"
	"strings"
	"time"
)

var (
	ErrTokenExpired           error = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh error = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         error = errors.New("请求令牌格式有误")
	ErrTokenInvalid           error = errors.New("请求令牌无效")
	ErrHeaderEmpty            error = errors.New("需要令牌才能访问")
	ErrHeaderMalformed        error = errors.New("请求头中 Authorization 格式有误")
)

// 定义一个JWT对象
type JWT struct {
	// 秘钥，用以加密 JWT，读取配置信息 app.key
	SignKey []byte
	// 刷新秘钥最大过期时间
	MaxRefresh time.Duration
}

// 自定义载荷
type JWTCustomClaims struct {
	UserID       string `json:"user_id"`
	UserName     string `json:"user_name"`
	ExpireAtTime int64  `json:"expire_time"`
	jwtpkg.StandardClaims
	// StandardClaims 结构体实现了 Claims 接口继承了  Valid() 方法
	// JWT 规定了7个官方字段，提供使用:
	// - iss (issuer)：发布者
	// - sub (subject)：主题
	// - iat (Issued At)：生成签名的时间
	// - exp (expiration time)：签名过期时间
	// - aud (audience)：观众，相当于接受者
	// - nbf (Not Before)：生效时间
	// - jti (JWT ID)：编号

}

// JWT初始化
func NewJWT() *JWT {
	return &JWT{
		SignKey:    []byte(config.GetString("app.key")),
		MaxRefresh: time.Duration(config.GetInt64("jwt.max_refresh_time")) * time.Minute,
	}
}

// createToken创建 Token，内部使用，外部请调用 IssueToken
func (jwt *JWT) createToken(claims JWTCustomClaims) (string, error) {
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodES256, claims)
	return token.SignedString(jwt.SignKey)
}

// IssueToken 生成 Token，外部调用,在登录成功时调用
func (jwt *JWT) IssueToken(userID, userName string) string {
	// 1. 构造用户 claims 信息
	expireAtTime := jwt.expireAtTime()
	claims := JWTCustomClaims{
		userID,
		userName,
		expireAtTime,
		jwtpkg.StandardClaims{
			NotBefore: app.TimenowInTimezone().Unix(),
			IssuedAt:  app.TimenowInTimezone().Unix(),
			ExpiresAt: expireAtTime,
			Issuer:    config.GetString("app.name"),
		},
	}
	// 2. 根据 claims 生成token对象
	token, err := jwt.createToken(claims)
	if err != nil {
		logger.LogIf(err)
		return ""
	}
	return token
}

// ParserToken 解析 Token，中间件中调用
func (jwt *JWT) ParserToken(c *gin.Context) (*JWTCustomClaims, error) {
	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return nil, parseErr
	}
	// 1. 调用 jwt 库解析用户传参的 Token
	token, err := jwt.parseTokenString(tokenString)
	// 2.解析是否出错
	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)
		if ok {
			//错误种类
			if validationErr.Errors == jwtpkg.ValidationErrorMalformed {
				return nil, ErrTokenMalformed
			} else if validationErr.Errors == jwtpkg.ValidationErrorExpired {
				return nil, ErrTokenExpired
			}

		}
		return nil, ErrTokenInvalid
	}
	// 3.解析后再数据结构校验
	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		//校验通过
		return claims, nil
	}
	return nil, ErrTokenInvalid
}

// 获取请求头token
func (jwt *JWT) getTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrHeaderEmpty
	}
	//使用空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2) && parts[0] == "Bearer" {
		return "", ErrHeaderEmpty
	}
	return parts[1], nil
}

// 使用 jwtpkg.ParseWithClaims 解析 Token
func (jwt *JWT) parseTokenString(tokenString string) (*jwtpkg.Token, error) {
	return jwtpkg.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwtpkg.Token) (interface{}, error) {
		return jwt.SignKey, nil
	})
}

// 获取过期时间
func (jwt *JWT) expireAtTime() int64 {
	timenow := app.TimenowInTimezone()

	var expireTime int64
	if app.IsLocal() {
		expireTime = config.GetInt64("jwt.debug_exprire_time")
	} else {
		expireTime = config.GetInt64("jwt.exprire_time")
	}
	expire := time.Duration(expireTime) * time.Minute
	return timenow.Add(expire).Unix()
}
