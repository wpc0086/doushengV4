package db

type Video struct {
	Id            int64  `json:"id,omitempty" gorm:"AUTO_INCREMENT;primary_key"`
	AuthorId      int64  `json:"author_id,omitempty" gorm:"index"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount uint   `json:"favorite_count,omitempty"`
	CommentCount  uint   `json:"comment_count,omitempty"`
	Title         string `json:"title,omitempty"`
	CreatedAt     int64  `json:"created_at" gorm:"index"`
}

func GetVideoListByUserId(userId int64) (videos []*Video, err error) {
	err = DB.Find(&videos, "author_id = ?", userId).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}
