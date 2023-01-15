package entity

type Series struct {
	ID              string   `json:"id"`              // UUID
	DisplayName     string   `json:"displayName"`     // 表示名
	Description     string   `json:"description"`     // 説明文
	ImageURL        string   `json:"imageUrl"`        // 画像URL
	CastIDs         []string `json:"castIds"`         // キャストIDs
	DisplayOrder    int32    `json:"displayOrder"`    // 表示順
	IsSingleEpisode bool     `json:"isSingleEpisode"` // シングルエピソードか
	EpisodeID       string   `json:"episodeId"`       // EpisodeID(SingleEpisodeがtrueのときのみ)
}

type SeriesMulti []*Series

type ListSeriesMultiResponse struct {
	SeriesMulti SeriesMulti `json:"series"`
	Casts       Casts       `json:"casts"`
}
