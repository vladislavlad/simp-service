package model

import "time"

type Comment struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	Text      string     `json:"text"`
	Author    string     `json:"author"`
	Target    string     `json:"target"`
	TargetId  uint       `json:"targetId"`
	ReplyTo   uint       `json:"replyTo"`
	CreatedAt time.Time  `gorm:"type:time;autoCreateTime:false" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"type:time;autoUpdateTime:false" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"type:time;autoUpdateTime:false" json:"deletedAt"`
}
