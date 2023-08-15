package models

type Video struct {
	ID            int    `json:"id"` // 视频的唯一ID，注意不是用户的
	FavoriteCount int    `json:"favorite_count"`
	CommentCount  int    `json:"comment_count"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	Title         string `json:"title"`
	IsFavorite    bool   `json:"is_favorite"`
	Author        `json:"author"`
}
type Author struct {
	ID              int64  `json:"id"`
	FollowCount     int    `json:"follow_count"`
	FollowerCount   int    `json:"follower_count"`
	WorkCount       int    `json:"work_count"`
	FavoriteCount   int    `json:"favorite_count"`
	Name            string `json:"name"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  string `json:"total_favorited"`
	IsFollow        bool   `json:"is_follow"`
}
