package model

import (
	"time"
)

type Comment struct {
	Id        uint      `gorm:"column:comment_id;primary_key"`
	TorrentId uint      `gorm:"column:torrent_id"`
	UserId    uint      `gorm:"column:user_id"`
	Content   string    `gorm:"column:content"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	Torrent   *Torrents `gorm:"ForeignKey:TorrentId"`
	User      *User     `gorm:"ForeignKey:UserId"`
}

type OldComment struct {
	TorrentId uint      `gorm:"column:torrent_id"`
	Username  string    `gorm:"column:username"`
	Content   string    `gorm:"column:content"`
	Date      time.Time `gorm:"column:date"`

	Torrent   *Torrents `gorm:"ForeignKey:TorrentId"`
}

func (c OldComment) TableName() string {
	// cba to renamed this in the db
	return "comments_old"
}
