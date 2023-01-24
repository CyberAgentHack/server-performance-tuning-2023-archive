package entity

type Genre struct {
	ID          string `json:"id"`          // UUID
	DisplayName string `json:"displayName"` // 表示名
}

type Genres []*Genre
