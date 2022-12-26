package entity

type Episode struct {
	ID           string  `json:"id"`                 // UUID
	DisplayName  string  `json:"displayName"`        // 表示名
	Description  string  `json:"description"`        // 説明
	ImageURL     string  `json:"imageUrl"`           // 画像URL
	SeriesID     string  `json:"seriesId"`           // シーズンID
	SeasonID     *string `json:"seasonId,omitempty"` // シーズンID
	DisplayOrder int32   `json:"displayOrder"`       // 表示順
}

type Episodes []*Episode

type ListEpisodesResponse struct {
	Episodes Episodes    `json:"episodes"`
	Series   SeriesMulti `json:"series"`
	Seasons  Seasons     `json:"seasons"`
}
