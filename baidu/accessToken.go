package baidu

type GetAcessTokenParam struct {
	GrantType    string
	Code         string
	ClientId     string
	SlientSecret string
}

type GetAcessTokenResponse struct {
	ExpiresIn     string `json:"expires_in"`
	RefreshToken  string `json:"refresh_token"`
	AccessToken   string `json:"access_token"`
	SessionSecret string `json:"session_secret"`
	SessionKey    string `json:"session_key"`
	Scope         string `json:"scope"`
}
