package model

type HideStrJson struct {
	HideStr string `json:"hide_str"`//需要隐藏的信息
	CarrierStr string `json:"carrier_str"`//载体字符串
	HideResultStr string `json:"hide_result_str"`//隐藏的信息插入载体字符串后的结果
	Optiontype string `json:"optiontype"` //1表示加密、2表示解密
}

type HideResult struct {
	HideStr string `json:"hide_str"`
	HideResultStr string `json:"hide_result_str"`
	HideResultLen int `json:"hide_result_len"`
}

type ClientInfo struct {
	UserAgent string `json:"user_agent"`
	RemoteIp string `json:"remote_ip"`

}

type GenPwd struct {
	PwdLength int64 `json:"pwd_length"`
	PwdType int64 `json:"pwd_type"`
}