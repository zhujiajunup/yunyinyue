package common

import "fmt"

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
	Nickname     string `json:"nickname"`
	UserType     int    `json:"userType"`
}

type Artist struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Alias     string `json:"alias"`
	AliasName string `json:"aliasName"`
}

type Album struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Company     string `json:"company"`
	PublishTime string `json:"publishTime"`
	Language    string `json:"language"`
}

type Song struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Artist Artist `json:"artist"`
	Album  Album  `json:"album"`
}

func (comment Comment) String() string {
	return fmt.Sprintf("%s: %s", comment.User.String(), comment.Content)
}

func (user YunyinyueUser) String() string {
	return user.Nickname
}
