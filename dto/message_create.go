package dto

import "github.com/JustAnotherID/botgo/dto/keyboard"

type MsgType int

const (
	MsgText MsgType = iota
	MsgPicAndText
	MsgMarkdown
	MsgArk
	MsgEmbed
	MsgInputNotify MsgType = 6
	MsgMedia       MsgType = 7
)

// MessageToCreate 发送消息结构体定义
type MessageToCreate struct {
	Content string  `json:"content,omitempty"`
	MsgType MsgType `json:"msg_type"`
	Embed   *Embed  `json:"embed,omitempty"`
	Ark     *Ark    `json:"ark,omitempty"`
	Image   string  `json:"image,omitempty"`
	Media   *Media  `json:"media,omitempty"`
	// 要回复的消息id，为空是主动消息，公域机器人会异步审核，不为空是被动消息，公域机器人会校验语料
	MsgID            string                    `json:"msg_id,omitempty"`
	MessageReference *MessageReference         `json:"message_reference,omitempty"`
	Markdown         *Markdown                 `json:"markdown,omitempty"`
	Keyboard         *keyboard.MessageKeyboard `json:"keyboard,omitempty"`     // 消息按钮组件
	EventID          string                    `json:"event_id,omitempty"`     // 要回复的事件id, 逻辑同MsgID
	MsgSeq           int                       `json:"msg_seq"`                // 回复消息的序号，与 msg_id 联合使用，避免相同消息 id 回复重复发送，不填默认是 1
	InputNotify      *InputNotify              `json:"input_notify,omitempty"` // 输入提醒，仅 msg_type = 6 时使用
}

// MessageReference 引用消息
type MessageReference struct {
	MessageID             string `json:"message_id"`               // 消息 id
	IgnoreGetMessageError bool   `json:"ignore_get_message_error"` // 是否忽律获取消息失败错误
}

// Markdown markdown 消息
type Markdown struct {
	TemplateID int               `json:"template_id"` // 模版 id
	Params     []*MarkdownParams `json:"params"`      // 模版参数
	Content    string            `json:"content"`     // 原生 markdown
}

// MarkdownParams markdown 模版参数 键值对
type MarkdownParams struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

// SettingGuideToCreate 发送引导消息的结构体
type SettingGuideToCreate struct {
	Content      string        `json:"content,omitempty"`       // 频道内发引导消息可以带@
	SettingGuide *SettingGuide `json:"setting_guide,omitempty"` // 设置引导
}

// SettingGuide 设置引导
type SettingGuide struct {
	// 频道ID, 当通过私信发送设置引导消息时，需要指定guild_id
	GuildID string `json:"guild_id"`
}

type InputType int

const (
	InputEntering   InputType = 1 // 对方正在输入
	InputCancel     InputType = 2 // 手动取消展示
	InputSpeaking   InputType = 3 // 对方正在讲话
	InputGenerating InputType = 4 // 正在生成
	InputImagining  InputType = 5 // 正在想象
)

// InputNotify 输入提醒
type InputNotify struct {
	InputType   InputType `json:"input_type"`
	InputSecond int       `json:"input_second"` // 持续秒数，1-60 之间
}
