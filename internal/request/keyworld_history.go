package request

type KeyworldHistoryInfo struct {
	ChatId             int64  `json:"chat_id" form:"chat_id" gorm:"column:chat_id;"`
	SenderId           int64  `json:"sender_id" form:"sender_id" gorm:"column:sender_id;"`
	SenderUsername     string `json:"sender_username" form:"sender_username" gorm:"column:sender_username;"`
	MessageId          int64  `json:"message_id" form:"message_id" gorm:"column:message_id;"`
	MessageLink        string `json:"message_link" form:"message_link" gorm:"column:message_link;"`
	MessageContentText string `json:"message_content_text" form:"message_content_text" gorm:"column:message_content_text;"`
	MessageDateEx      string `json:"message_date_ex" form:"message_date_ex" gorm:"column:message_date_ex;"`
	MessageDate        int64  `json:"message_date" form:"message_date" gorm:"column:message_date;"`

	KeyWorld string `json:"key_world" form:"key_world" gorm:"column:key_world;"`

	// 核验请求id
	ById string `json:"by_id" form:"by_id" gorm:"-"`

	Page   int    `json:"-" form:"page" gorm:"-"`
	Size   int    `json:"-" form:"size" gorm:"-"`
	Filter string `json:"-" form:"filter" gorm:"-"`
}
