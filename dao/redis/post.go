package redis

import (
	"strconv"
	"time"
	"web_app/models"

	"github.com/go-redis/redis"
)

func getIDsFormKey(key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := start + size - 1
	return client.ZRevRange(key, start, end).Result()
}

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}

	return getIDsFormKey(key, p.Page, p.Size)
}

func GetPostVoteData(ids []string) (data []int64, err error) {
	//data = make([]int64, 0, len(ids))
	//for _, id := range ids{
	//	key := getRedisKey(KeyPostVotedZSetPF+id)
	//	v := client.ZCount(key, "1", "1").Val()
	//	data = append(data, v)
	//}

	pipeline := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPF + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}

func GetCommunityPostIDsInOrder(p *models.ParamPostList) ([]string, error) {

	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}
	communityKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(p.CommunityID)))

	key := orderKey + strconv.Itoa(int(p.CommunityID))
	if client.Exists(key).Val() < 1 {
		pipeline := client.Pipeline()

		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, communityKey, orderKey)
		pipeline.Expire(key, 60*time.Second)
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}

	return getIDsFormKey(key, p.Page, p.Size)
}
