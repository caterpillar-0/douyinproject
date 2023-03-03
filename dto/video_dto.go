package dto

type VideoDTO struct {
	ID            int64       `json:"id"`             // 视频唯一标识
	Author        UserInfoDTO `json:"author"`         // 视频作者信息
	PlayURL       string      `json:"play_url"`       // 视频播放地址
	CoverURL      string      `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64       `json:"favorite_count"` // 视频的点赞总数
	CommentCount  int64       `json:"comment_count"`  // 视频的评论总数
	IsFavorite    bool        `json:"is_favorite"`    // true-已点赞，false-未点赞
	Title         string      `json:"title"`          // 视频标题
}

type VideoResponse struct {
	Response
	VideoList []*VideoDTO `json:"video_list"`
	NextTime  int64       `json:"next_time,omitempty"`
}
