package hw_push

const (
	PUSH  = "https://api.push.hicloud.com/pushsend.do"
	TOKEN = "https://login.cloud.huawei.com/oauth2/v2/token"
)

type ext struct {
	Url     string `json:"url"`
	Func    string `json:"func"`
	Title   string `json:"title"`
	Action  string `json:"action"`
	Collect string `json:"collect"`
	Content string `json:"content"`
}

type hps struct {
	Msg msg `json:"msg"`
	Ext ext `json:"ext"`
}

type msg struct {
	Type   int    `json:"type"`
	Body   body   `json:"body"`
	Action action `json:"action"`
}

type vers struct {
	Ver   string `json:"ver"`
	AppID string `json:"appId"`
}

type body struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type param struct {
	Intent     string `json:"intent"`
	AppPkgName string `json:"appPkgName"`
}

type action struct {
	Type  int   `json:"type"`
	Param param `json:"param"`
}

type extObj struct {
	Name string
}

type tokenResponse struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
}
