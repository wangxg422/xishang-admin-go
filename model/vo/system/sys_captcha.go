package system

type SysCaptchaVO struct {
	CaptchaEnabled bool   `json:"captchaEnabled"`
	CaptchaId      string `json:"captchaId"`
	PicPath        string `json:"picPath"`
	CaptchaLength  int    `json:"captchaLength"`
}
