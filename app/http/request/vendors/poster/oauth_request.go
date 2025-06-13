package vendors_poster_request

type OauthRequest struct {
	Code    string `form:"code"`
	Account string `form:"account"`
}
