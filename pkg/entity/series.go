package entity

type Series struct {
	ID          string `json:"id"`          // UUID
	DisplayName string `json:"displayName"` // 表示名
	Description string `json:"description"` // 説明文
	ImageURL    string `json:"imageUrl"`    // 画像URL
	GenreID     string `json:"genreId"`     // ジャンルID
}

type SeriesMulti []*Series

type ListSeriesMultiResponse struct {
	SeriesMulti SeriesMulti `json:"series"`
	Genres      Genres      `json:"genres"`
}

func (s SeriesMulti) GenreIDs() []string {
	ret := make([]string, 0, len(s))
	for i := range s {
		ret = append(ret, s[i].GenreID)
	}
	return ret
}
