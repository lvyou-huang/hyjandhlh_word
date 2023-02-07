package model

type Accountinfo struct {
	Phone        int    `json:"phone"`
	Weixin       string `json:"weixin"`
	Xinlang      string `json:"xinlang"`
	Github       string `json:"github"`
	Phoneoremail string `json:"phoneoremail"`
}
