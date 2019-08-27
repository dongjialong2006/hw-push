package hw_push

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	id     string
	ctx    string
	secret string
}

func NewClient(clientID string, clientSecret string) *Client {
	data, _ := json.Marshal(&vers{
		Ver:   "1",
		AppID: clientID,
	})

	return &Client{
		id:     clientID,
		ctx:    string(data),
		secret: clientSecret,
	}
}

func (c *Client) serve(url string, data url.Values) ([]byte, error) {
	u := ioutil.NopCloser(strings.NewReader(data.Encode()))
	r, err := http.Post(url, "application/x-www-form-urlencoded", u)
	if err != nil {

		return []byte(""), err
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {

		return []byte(""), err
	}
	return b, err
}

func (c *Client) GetToken() (string, error) {
	values := make(url.Values)
	values["grant_type"] = []string{"client_credentials"}
	values["client_id"] = []string{c.id}
	values["client_secret"] = []string{c.secret}

	res, err := c.serve(TOKEN, values)
	if nil != err {
		return "", err
	}
	var resp = &tokenResponse{}
	err = json.Unmarshal(res, resp)
	if err != nil {
		return "", err
	}
	return resp.AccessToken, nil
}

func (c *Client) Push(tokens []string, payload string) (string, error) {
	token, err := c.GetToken()
	if nil != err {
		return "", err
	}

	data, _ := json.Marshal(tokens)

	var originParam = map[string]string{
		"nsp_ts":            strconv.Itoa(int(time.Now().Unix())),
		"nsp_svc":           "openpush.message.api.send",
		"payload":           payload,
		"expire_time":       time.Now().Format("2006-01-02T15:04"),
		"access_token":      token,
		"device_token_list": string(data),
	}

	param := make(url.Values)
	param["nsp_ts"] = []string{originParam["nsp_ts"]}
	param["payload"] = []string{originParam["payload"]}
	param["nsp_svc"] = []string{originParam["nsp_svc"]}
	param["access_token"] = []string{originParam["access_token"]}
	param["device_token_list"] = []string{originParam["device_token_list"]}

	res, err := c.serve(fmt.Sprintf("%s?nsp_ctx=%s", PUSH, url.QueryEscape(c.ctx)), param)

	return string(res), err
}
