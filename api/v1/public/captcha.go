package public

import (
	"backend/common/response"
	"backend/global"
	"backend/initial/logger"
	"backend/model/vo"
	"backend/utils/captcha"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type CaptchaApi struct {
}

func (m *CaptchaApi) GenCaptcha(c *gin.Context) {

	config := global.AppConfig.Captcha
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(config.ImgHeight, config.ImgWidth, config.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, captcha.GetRedisStore().UseWithCtx(c))

	id, b64s, err := cp.Generate()
	if err != nil {
		logger.Error("验证码获取失败!", zap.Error(err))
		response.FailWithMessage("验证码获取失败", c)
		return
	}
	response.OkWithDetailed(vo.SysCaptchaVO{
		CaptchaEnabled: true,
		CaptchaId:      id,
		PicPath:        b64s,
		CaptchaLength:  config.KeyLong,
	}, "验证码获取成功", c)
}
