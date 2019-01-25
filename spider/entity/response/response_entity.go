package response

import "yunyinyue/spider/entity/common"

type CommentResp struct {
	IsMusician bool             `json:"isMusician"`
	UserId     int32            `json:"userId"`
	MoreHot    bool             `json:"moreHot"`
	Comments   []common.Comment `json:"comments"`
	Total      int32            `json:"total"`
}

type SongDetail struct {
	Song common.Song `json:"song"`
	// ignore other
}
type PlayRecord struct {
	PlayCount  int        `json:"playCount"`
	Score      int        `json:"score"`
	SongDetail SongDetail `json:"song"`
}

type PlayRecordResp struct {
	Code     int          `json:"code"`
	AllData  []PlayRecord `json:"allData"`
	WeekData []PlayRecord `json:"weekData"`
}
