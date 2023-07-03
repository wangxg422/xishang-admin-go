package system

type SysLoginDTO struct {
	UserName    string `json:"username,omitempty" binding:"required"`
	Password    string `json:"password,omitempty" binding:"required"`
	CaptchaCode string `json:"captchaCode,omitempty"`
	CaptchaId   string `json:"captchaId,omitempty"`
}
