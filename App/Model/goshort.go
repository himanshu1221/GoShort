package model

func GetAllShorts() ([]Goshort, error) {
	var goshorts []Goshort

	tx := db.Find(&goshorts)
	if tx.Error != nil {
		return []Goshort{}, tx.Error
	}

	return goshorts, nil
}

func SingleShorts(id uint64) (Goshort, error) {
	var goshort Goshort

	tx := db.Where("id=?", id).First(&goshort)

	if tx.Error != nil {
		return goshort, tx.Error
	}

	return goshort, nil
}

func CreateShort(goshort Goshort) error {
	tx := db.Create(&goshort)
	return tx.Error
}

func UpdateShort(goshort Goshort) error {
	tx := db.Save(&goshort)
	return tx.Error
}

func DeleteShort(id uint64) error {
	tx := db.Unscoped().Delete(&Goshort{}, id)
	return tx.Error
}

func FindShortByUrl(url string) (Goshort, error) {
	var goshort Goshort
	tx := db.Where("goshort=?", url).First(&goshort)
	return goshort, tx.Error
}
