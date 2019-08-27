package hw_push

import (
	"encoding/json"
)

type message struct {
	HPS hps `json:"hps"`
}

func NewMessage() *message {
	return &message{
		HPS: hps{
			Msg: msg{
				Type: 3, //1, 透传异步消息; 3, 系统通知栏异步消息;
				Body: body{},
				Action: action{
					Type: 1, //1, 自定义行为; 2, 打开URL; 3, 打开App;
					Param: param{
						Intent:     "#Intent;compo=com.rvr/.Activity;S.W=U;end",
						AppPkgName: "",
					},
				},
			},
			Ext: ext{},
		},
	}
}

func (m *message) SetContent(content string) {
	m.HPS.Msg.Body.Content = content
}

func (m *message) SetTitle(title string) {
	m.HPS.Msg.Body.Title = title
}

func (m *message) SetIntent(intent string) {
	m.HPS.Msg.Action.Param.Intent = intent
}

func (m *message) SetAppPkgName(appPkgName string) {
	m.HPS.Msg.Action.Param.AppPkgName = appPkgName
}

func (m *message) SetExtAction(Action string) {
	m.HPS.Ext.Action = Action
}
func (m *message) SetExtFunc(Func string) {
	m.HPS.Ext.Func = Func
}
func (m *message) SetExtCollect(collect string) {
	m.HPS.Ext.Collect = collect
}
func (m *message) SetExtTitle(title string) {
	m.HPS.Ext.Title = title
}
func (m *message) SetExtContent(content string) {
	m.HPS.Ext.Collect = content
}

func (m *message) SetExtUrl(url string) {
	m.HPS.Ext.Url = url
}

func (m *message) Json() string {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
