package entity

type Cast struct {
	ID          string `json:"id"`          // UUID
	DisplayName string `json:"displayName"` // 表示名
	ImageURL    string `json:"imageUrl"`    // 画像URL
	Biography   string `json:"biography"`   // 自己紹介文
	KanaName    string `json:"kanaName"`    // かな文字の名前
}

type Casts []*Cast
