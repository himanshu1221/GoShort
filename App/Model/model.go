package model

type GoShort struct {
	ID       uint64 `json:"id" gorm:"primaryKey"` //Go Object Relational Mapping,
	Redirect string `json:"redirect" gorm:"not null"`
	GoShort  string `json:"goshort" gorm:"unique;not null"`
	Clicked  uint64 `json:"clicked"`
	Random   bool   `json:"random"`
}

func Setup() {

}
