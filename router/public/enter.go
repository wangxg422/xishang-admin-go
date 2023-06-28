package public

import "backend/api/v1/public"

type PublicRouterGroup struct {
	PublicRouter
}

var PublicRouterGroupApp = new(PublicRouterGroup)

var (
	loginApi   = public.LoginApi{}
	captchaApi = public.CaptchaApi{}
)
