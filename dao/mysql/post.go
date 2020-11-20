package mysql

import "web_app/models"

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
				post_id, title, content, author_id, community_id)
				values(?,?,?,?,?)
				`
	_, err = db.Exec(sqlStr, p.PostID, p.Title, p.Content, p.AuthorId, p.CommunityID)
	return
}
