package mysql

import (
	"bluebell_backend/models"

	"go.uber.org/zap"
)

func CreateComment(comment *models.Comment) (err error) {
	sqlStr := `insert into comment(
	comment_id, content, post_id, author_id, author_name)
	values(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, comment.CommentID, comment.Content, comment.PostID,
		comment.AuthorID, comment.AuthorName)
	if err != nil {
		zap.L().Error("insert comment failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}

func GetCommentListByIDs(ids string) (commentList []*models.Comment, err error) {
	sqlStr := `select comment_id, content, post_id, author_id, author_name, create_time
	from comment
	where post_id = ?`
	err = db.Select(&commentList, sqlStr, ids)
	return
}
