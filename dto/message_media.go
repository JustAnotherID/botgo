package dto

// MessageMediaToCreate 创建富媒体消息
type MessageMediaToCreate struct {
	FileType   int    `json:"file_type"`
	URL        string `json:"url"`
	SrvSendMsg bool   `json:"srv_send_msg"`
	FileData   []byte `json:"file_data"`
}

type Media struct {
	FileUUID string `json:"file_uuid"`
	FileInfo string `json:"file_info"`
	Ttl      int    `json:"ttl"`
}

type MediaMessage struct {
	Message
	Media
}
