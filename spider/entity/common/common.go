package common

import "fmt"

type Comment struct {
	User    YunyinyueUser `json:"user"`
	Time    uint64        `json:"time"`
	Content string        `json:"content"`
	SongId  string        `json:"-"`
}

type YunyinyueUser struct {
	LocationInfo string `json:"locationInfo" db:"location_info"`
	ExpertTags   string `json:"expertTags" db:"expert_tags"`
	AvatarUrl    string `json:"avatarUrl" db:"avatar_url"`
	UserId       int32  `json:"userId" db:"user_id"`
	Nickname     string `json:"nickname" db:"nickname"`
	UserType     int    `json:"userType" db:"user_type"`
}

type Artist struct {
	Id        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Alias     string `json:"alias" db:"alias"`
	AliasName string `json:"aliasName" db:"alias_name"`
}

type Album struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Company     string `json:"company"`
	PublishTime string `json:"publishTime"`
	Language    string `json:"language"`
}

type Song struct {
	Id     string `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Artist Artist `json:"artist" `
	Album  Album  `json:"album"`
}

func (comment Comment) String() string {
	return fmt.Sprintf("%s: %s", comment.User.String(), comment.Content)
}

func (user YunyinyueUser) String() string {
	return user.Nickname
}
