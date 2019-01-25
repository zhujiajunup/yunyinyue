package entity

import "fmt"

type CommentResp struct {
	IsMusician bool       `json:"isMusician"`
	UserId     int32      `json:"userId"`
	MoreHot    bool       `json:"moreHot"`
	Comments   [] Comment `json:"comments"`
	Total      int32      `json:"total"`
}

type Comment struct {
	User    YunyinyueUser `json:"user"`
	Time    uint64        `json:"time"`
	Content string        `json:"content"`
}

type YunyinyueUser struct {
	LocationInfo string `json:"locationInfo"`
	ExpertTags   string `json:"expertTags"`
	AvatarUrl    string `json:"avatarUrl"`
	UserId       int32  `json:"userId"`
	Nickname     string  `json:"nickname"`
	UserType     int    `json:"userType"`
}

// comment request body
type CommentRequestBody struct {
	Rid       string `json:"rid"`
	Offset    string `json:"offset"`
	Totail    string `json:"totail"`
	Limit     string `json:"limit"`
	CsrfToken string `json:"csrf_token"`
}

func (comment Comment) String() string{
	return fmt.Sprintf("%s: %s", comment.User.String(), comment.Content)
}

func (user YunyinyueUser) String() string{
	return user.Nickname
}
