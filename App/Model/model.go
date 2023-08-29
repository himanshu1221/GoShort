package model

type Goshort struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Redirect string `json:"redirect"`
	Goshort  string `json:"goshort" gorm:"unique;not null"`
	Clicked  uint64 `json:"clicked"`
	Random   bool   `json:"random"`
}
