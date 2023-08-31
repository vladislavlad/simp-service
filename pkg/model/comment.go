package model

type Comment struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Text     string `json:"text"`
	Author   string `json:"author"`
	Target   string `json:"target"`
	TargetId uint   `json:"targetId"`
	ReplyTo  uint   `json:"replyTo"`
}
