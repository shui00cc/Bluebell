package models

import "time"

type Comment struct {
	PostID     uint64    `db:"post_id" json:"post_id,string"`
	CommentID  uint64    `db:"comment_id" json:"comment_id,string"`
	AuthorID   uint64    `db:"author_id" json:"author_id,string"`
	Content    string    `db:"content" json:"content"`
	AuthorName string    `db:"author_name" json:"author_name"`
	CreateTime time.Time `db:"create_time" json:"create_time"`
}
