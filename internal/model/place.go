package model

import (
	"leizhenpeng/go-mux-gorm-crud/internal/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Place struct {
	gorm.Model
	Name string `json:"name"`
	Star int    `json:"star"`
}

func InitDB() {
	db = config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Place{})
}

func GetAllPlaces() ([]Place, error) {
	var places []Place
	err := db.Find(&places).Error
	if err != nil {
		return nil, err
	}
	return places, nil
}

func GetPlaceByID(id int) (Place, error) {
	var place Place
	err := db.Where("id = ?", id).First(&place).Error
	if err != nil {
		return place, err
	}
	return place, nil
}

func (p *Place) CreatePlace() error {
	err := db.Create(&p).Error
	if err != nil {
		return err
	}
	return nil
}

func CreatePlace(place Place) (Place, error) {
	err := db.Create(&place).Error
	if err != nil {
		return place, err
	}
	return place, nil
}

func UpdatePlace(id int, place Place) (Place, error) {
	var p Place
	err := db.Model(&p).Where("id = ?", id).Updates(place).Error
	if err != nil {
		return place, err
	}

	//return updated place
	err2 := db.Where("id = ?", id).First(&place).Error
	if err2 != nil {
		return place, err2
	}
	return place, nil
}

func DeletePlace(id int) error {
	var place Place
	err := db.Where("id = ?", id).Delete(&place).Error
	if err != nil {
		return err
	}
	return nil
}
