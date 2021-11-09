package models

import "golang-rest-api/database"

func GetNilaiSamaV2(id_provinsi, id_rs, limit string) []Properties {

	// var DataNilaiObj Properties
	var DataNilaiArrObj []Properties

	// con := database.CreateCon()
	db, _ := database.KonekMysql()

	sqlQuery := `select 
					distinct pas_m.desano, pas_m.nama_desa, pas_m.id_rs
				from 
					pasien_maps pas_m
				join 
					kelurahan_maps kel_m on kel_m.desano = pas_m.desano
				where 
					pas_m.id_provinsi = ` + id_provinsi + `
				and 
					pas_m.id_rs = ` + id_rs + `
				and
					kel_m.id_rs = ` + id_rs + `
				and
					kel_m.id_provinsi = ` + id_provinsi + `
				order by pas_m.desano asc limit ` + limit
				
	db.Raw(sqlQuery).Scan(&DataNilaiArrObj)

	return DataNilaiArrObj
}