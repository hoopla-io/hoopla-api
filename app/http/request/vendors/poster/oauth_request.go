package vendors_poster_request

type OauthRequest struct {
	Code    string `query:"code"`
	Account string `query:"account"`
}
