package logic

import (
	"strconv"
	"web_app/dao/redis"
	"web_app/models"

	"go.uber.org/zap"
)

// vote function
// 1. user vote data
// 2. vote status

// vote recommendation algorithm: http://www.ruanyifeng.com/blog/algorithm/

// one vote 432 86400/200 -> 200 put on first

/* situations

direction=1:
	[-1, 0 ,1] : [-1, 0, 1]

limitation:
	can only vote in one week, one week later, sync to mysql
	KeyPostVotedZSetPF
*/

func VoteForPost(userID int64, p *models.ParamVoteData) error {
	// 1. limitation
	// 2. update score
	// 3. record data
	zap.L().Debug("VoteForPost",
		zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction),
	)
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))

}
