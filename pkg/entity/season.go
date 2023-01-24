package entity

type Season struct {
	ID           string   `json:"id"`           // UUID
	DisplayName  string   `json:"displayName"`  // 表示名
	ImageURL     string   `json:"imageUrl"`     // 画像URL
	GenreIDs     []string `json:"genreIds"`     // ジャンルIDs
	SeriesID     string   `json:"seriesId"`     // VideoSeriesID
	DisplayOrder int32    `json:"displayOrder"` // リスト時の表示順
}

type Seasons []*Season

type ListSeasonsResponse struct {
	Seasons Seasons `json:"seasons"`
	Genres  Genres  `json:"genres"`
}

func (s Seasons) GenreIDs() []string {
	ret := make([]string, 0, len(s))
	for i := range ret {
		ret = append(ret, s[i].GenreIDs...)
	}
	return ret
}
