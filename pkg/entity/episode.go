package entity

import "time"

type Episode struct {
	ID               string    `json:"id"`               // UUID
	DisplayName      string    `json:"displayName"`      // 表示名
	Description      string    `json:"description"`      // 説明
	ImageURL         string    `json:"imageUrl"`         // 画像URL
	CastIDs          []string  `json:"castIds"`          // キャストIDs
	SeasonID         string    `json:"seasonId"`         // シーズンID
	PublishStartTime time.Time `json:"publishStartTime"` // 公開開始時刻
	DisplayOrder     int32     `json:"displayOrder"`     // 表示順
}

type Episodes []*Episode

type ListEpisodesResponse struct {
	Episodes Episodes `json:"episodes"`
	Casts    Casts    `json:"casts"`
}

func (e Episodes) CastIDs() []string {
	ret := make([]string, 0, len(e))
	for i := range ret {
		ret = append(ret, e[i].CastIDs...)
	}
	return ret
}
