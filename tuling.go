package main

// Tuling 图灵
type Tuling struct {
	URL  string           `yaml:"url"`
	Keys map[string]Rebot `yaml:"keys"`
}

// Rebot 机器人
type Rebot struct {
	Name string `yaml:"Name"`
	Key  string `yaml:"key"`
}

// UserInfo 用户参数
type UserInfo struct {
	APIKey string `json:"apiKey"`
	UserID string `json:"userId"`
}

// Perception 输入信息
type Perception struct {
	InputText InputText `json:"inputText"`
}

// InputText 文本信息
type InputText struct {
	Text string `json:"text"`
}

// Reply 回复数据格式
type Reply struct {
	Intent  ReplyIntent   `json:"intent"`
	Results []ReplyResult `json:"results"`
}

// ReplyIntent 图灵请求意图
type ReplyIntent struct {
	Code       int    `json:"code"`       // 输出功能code
	IntentName string `json:"intentName"` // 意图名称
}

// ReplyResult 输出结果集
type ReplyResult struct {
	ResultType string      `json:"resultType"` // 输入类型
	Values     ReplyValues `json:"values"`     // 输出值
}

// ReplyValues 输出值
type ReplyValues struct {
	Text string `json:"text"`
	URL  string `json:"url"`
}
