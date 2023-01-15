package entity

type Season struct {
	ID           string   `json:"id"`           // UUID
	DisplayName  string   `json:"displayName"`  // 表示名
	ImageURL     string   `json:"imageUrl"`     // 画像URL
	CastIDs      []string `json:"castIds"`      // キャストIDs
	SeriesID     string   `json:"seriesId"`     // VideoSeriesID
	DisplayOrder int32    `json:"displayOrder"` // リスト時の表示順
}

type Seasons []*Season
