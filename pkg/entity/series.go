package entity

type Series struct {
	ID string `json:"id"`
}

type SeriesMulti []*Series
