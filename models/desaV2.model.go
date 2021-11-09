package models

import "golang-rest-api/database"

func GetLatLongDesanoDataV3(desano, id_provinsi, id_rs string) []GetLatLongDesano {
	// array
	var arrobjgetLatLongDesano []GetLatLongDesano

	db, _ := database.KonekMysql()

	//getGroupByDesano
	sqlQuery := `select 
					latitude, longitude 
				from 
					kelurahan_maps 
				where 
					desano= '` + desano + `' 
				and 
					id_rs= ` + id_rs + `
				and 
					id_provinsi= ` + id_provinsi

	db.Raw(sqlQuery).Scan(&arrobjgetLatLongDesano)
	
	return arrobjgetLatLongDesano
}