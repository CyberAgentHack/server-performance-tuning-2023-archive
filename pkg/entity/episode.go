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
