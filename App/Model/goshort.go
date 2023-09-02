package model

func GetAllShorts() ([]Goshort, error) {
	var goshorts []Goshort

	tx := db.Find(&goshorts)
	if tx.Error != nil {
		return []Goshort{}, tx.Error
	}

	return goshorts, nil
}
