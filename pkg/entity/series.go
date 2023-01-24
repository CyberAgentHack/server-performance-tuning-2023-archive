package entity

type Series struct {
	ID              string   `json:"id"`              // UUID
	DisplayName     string   `json:"displayName"`     // 表示名
	Description     string   `json:"description"`     // 説明文
	ImageURL        string   `json:"imageUrl"`        // 画像URL
	GenreIDs        []string `json:"genreIds"`        // ジャンルIDs
	DisplayOrder    int32    `json:"displayOrder"`    // 表示順
	IsSingleEpisode bool     `json:"isSingleEpisode"` // シングルエピソードか
	EpisodeID       string   `json:"episodeId"`       // EpisodeID(SingleEpisodeがtrueのときのみ)
}

type SeriesMulti []*Series

type ListSeriesMultiResponse struct {
	SeriesMulti SeriesMulti `json:"series"`
	Genres      Genres      `json:"genres"`
}

func (s SeriesMulti) GenreIDs() []string {
	ret := make([]string, 0, len(s))
	for i := range ret {
		ret = append(ret, s[i].GenreIDs...)
	}
	return ret
}
