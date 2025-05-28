package models

import "github.com/seccret404/Ecomerce-Api/config"

func MigateALL() {
	err := config.DB.AutoMigrate(
		&Product{},
		&CartItem{},
	)

	if err != nil{
		panic("gagal migrate" + err.Error())
	}

}
