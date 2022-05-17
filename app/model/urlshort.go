package model

func GetAllShorts() ([]URLShort, error) {
	var shorts []URLShort

	// transaction
	tx := db.Find(&shorts)
	if tx.Error != nil {
		return []URLShort{}, tx.Error
	}

	return shorts, nil
}

func GetShort(id uint64) (URLShort, error) {
	var short URLShort

	tx := db.Where("id = ?", id).First(&short)
	if tx.Error != nil {
		return URLShort{}, tx.Error
	}

	return short, nil
}

func CreateShort(short URLShort) error {
	tx := db.Create(&short)
	return tx.Error
}

func UpdateShort(short URLShort) error {
	tx := db.Save(&short)
	return tx.Error
}

func DeleteShort(id uint64) error {
	tx := db.Unscoped().Delete(&URLShort{}, id)
	return tx.Error
}

func FindByShortUrl(url string) (URLShort, error) {
	var short URLShort

	tx := db.Where("url_short = ?", url).First(&short)
	return short, tx.Error
}
