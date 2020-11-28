package logic

import (
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/models"
	"web_app/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	// 1.generate ID
	p.PostID = snowflake.GenID()

	err = mysql.CreatePost(p)

	if err != nil {
		return err
	}

	err = redis.CreatePost(p.PostID, p.CommunityID)
	return
}

func GetPostById(pid int64) (data *models.ApiPostDetail, err error) {
	// check data and interfaces
	post, err := mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostById(pid) failed", zap.Int64("pid", pid), zap.Error(err))
		return
	}
	user, err := mysql.GetUserById(post.AuthorId)
	if err != nil {
		zap.L().Error("mysql.GetUserById(post.AuthorId) failed",
			zap.Int64("author", post.AuthorId),
			zap.Error(err))
		return
	}

	communityDetail, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed",
			zap.Int64("community_id", post.CommunityID),
			zap.Error(err))
	}

	data = &models.ApiPostDetail{
		AutherName:      user.Username,
		Post:            post,
		CommunityDetail: communityDetail,
	}

	return
}

func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}

	data = make([]*models.ApiPostDetail, 0, len(posts))

	for _, post := range posts {
		user, err := mysql.GetUserById(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorId) failed",
				zap.Int64("author", post.AuthorId),
				zap.Error(err))
			continue
		}

		communityDetail, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err))
			continue
		}
		postdetail := &models.ApiPostDetail{
			AutherName:      user.Username,
			Post:            post,
			CommunityDetail: communityDetail,
		}
		data = append(data, postdetail)
	}
	return
}

func GetPostList2(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 len")
		return
	}

	zap.L().Debug("GetPostList2", zap.Any("ids", ids))

	posts, err := mysql.GetPostListByIDs(ids)

	if err != nil {
		return
	}

	zap.L().Debug("posts content", zap.Any("posts", posts))

	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}

	for idx, post := range posts {
		user, err := mysql.GetUserById(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorId) failed",
				zap.Int64("author", post.AuthorId),
				zap.Error(err))
			continue
		}

		communityDetail, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err))
			continue
		}
		postdetail := &models.ApiPostDetail{
			AutherName:      user.Username,
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: communityDetail,
		}
		data = append(data, postdetail)
	}
	return
}

func GetCommunityPostList(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {

	ids, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 len")
		return
	}

	zap.L().Debug("GetPostList2", zap.Any("ids", ids))

	posts, err := mysql.GetPostListByIDs(ids)

	if err != nil {
		return
	}

	zap.L().Debug("posts content", zap.Any("posts", posts))

	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}

	for idx, post := range posts {
		user, err := mysql.GetUserById(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorId) failed",
				zap.Int64("author", post.AuthorId),
				zap.Error(err))
			continue
		}

		communityDetail, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err))
			continue
		}
		postdetail := &models.ApiPostDetail{
			AutherName:      user.Username,
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: communityDetail,
		}
		data = append(data, postdetail)
	}
	return
}

func GetPostListNew(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	if p.CommunityID == 0 {
		data, err = GetPostList2(p)
	} else {
		data, err = GetCommunityPostList(p)
	}
	if err != nil {
		zap.L().Error("GetPostListNew failed", zap.Error(err))
		return nil, err
	}
	return
}
