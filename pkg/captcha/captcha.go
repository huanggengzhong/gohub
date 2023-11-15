package captcha

import (
	"github.com/mojocn/base64Captcha"
	"gohub/pkg/config"
	"gohub/pkg/redis"
	"sync"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

var once sync.Once

// 内部使用的 Captcha 对象
var internalCaptcha *Captcha

// 单例模式
func NewCaptcha() *Captcha {
	once.Do(func() {
		internalCaptcha = &Captcha{}

		drivery := base64Captcha.NewDriverDigit(
			config.GetInt("captcha.height"),      // 宽
			config.GetInt("captcha.width"),       // 高
			config.GetInt("captcha.length"),      // 长度
			config.GetFloat64("captcha.maxskew"), // 数字的最大倾斜角度
			config.GetInt("captcha.dotcount"),    // 图片背景里的混淆点数量
		)
		// 使用全局 Redis 对象，并配置存储 Key 的前缀
		store := RedisStore{
			RedisClient: redis.Redis,
			KeyPrefix:   config.GetString("app.name") + ":captcha",
		}

		// 实例化 base64Captcha 并赋值给内部使用的 internalCaptcha 对象
		internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(drivery, &store)
	})
	return internalCaptcha
}

// 生成图片验证码
func (c *Captcha) GenerateCaptcha() (id string, b64s string, err error) {
	return c.Base64Captcha.Generate()
}

// 校验图片验证码
func (c *Captcha) VerifyCaptcha(id string, anser string) (match bool) {
	// 方便本地和 API 自动测试
	//if !app.IsProduction() && id == config.GetString("captcha.testing_key") {
	//	return true
	//}
	// 第三个参数是验证后是否删除，我们选择 false
	// 这样方便用户多次提交，防止表单提交错误需要多次输入图片验证码
	return c.Base64Captcha.Verify(id, anser, false)
}
