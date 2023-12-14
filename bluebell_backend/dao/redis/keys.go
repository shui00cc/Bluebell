package redis

// redis key 注意使用命名空间的方式，方便查询和拆分
const (
	KeyPostInfoHashPrefix = "bluebell-plus:post:"      // Hash; bluebell-plus:post:$post_id——key——value; 存储帖子相关信息
	KeyPostTimeZSet       = "bluebell-plus:post:time"  // ZSet; bluebell-plus:post:time——Time——PostID; 帖子和发帖时间
	KeyPostScoreZSet      = "bluebell-plus:post:score" // ZSet; bluebell-plus:post:score——Score——PostID; 帖子及投票分数

	KeyPostVotedZSetPrefix    = "bluebell-plus:post:voted:" // ZSet; bluebell-plus:post:voted:$post_id——VoteNum——VoteUserID; 每个帖子投票值和投票人
	KeyCommunityPostSetPrefix = "bluebell-plus:community:"  // Set; bluebell-plus:community:$community_id——[PostIDs]; 社区下所有的帖子
)
