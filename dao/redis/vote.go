package redis

import (
	"errors"
	"strconv"
	"time"

	"go.uber.org/zap"

	"github.com/go-redis/redis"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVot      = 432
)

var (
	ErrVoteTimeExpire = errors.New("INVALID VOTE TIME")
	ErrVoteRepeat     = errors.New("REPEAT VOTE NOT ALLOWED")
)

func CreatePost(postID, communityID int64) error {

	pipeline := client.TxPipeline()

	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	}).Result()

	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	}).Result()

	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(communityID)))
	pipeline.SAdd(cKey, postID)
	_, err := pipeline.Exec()
	return err
}

func VoteForPost(userID, postID string, value float64) error {
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	zap.L().Info("print redis post time", zap.Any("post time: ", postTime))
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	ov := client.ZScore(getRedisKey(KeyPostVotedZSetPF+postID), userID).Val()
	if value == ov {
		return ErrVoteRepeat
	}

	diff := value - ov

	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), diff*scorePerVot, postID)

	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostScoreZSet+postID), userID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPF+postID), redis.Z{
			Score:  value,
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err
}
