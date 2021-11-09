package models

import (
	"fmt"
	"golang-rest-api/database"
	"net/http"
)

type GetGroupByDesano struct {
	Desano string `json:"desano"`
}

type GetLatLongDesano struct {
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}


type GetPasienMapIcd struct {
	Desano string `json:"desano"`
	Nama_desa string `json:"nama_desa"`
	Id_rs string `json:"id_rs"`
	Waktu string `json:"waktu"`
}

type GetPenyakitByKelurahan struct {
	Nama_penyakit string `json:"nama_penyakit"`
	Kode string `json:"kode"`
	Kode_icd string `json:"kode_icd"`
	Id_rs string `json:"id_rs"`
	Kode_kelurahan string `json:"kode_kelurahan"`
	Jumlah_Pasien string `json:"jml_pasien"`
}

func GetPasienMapIcdData(id_provinsi, id_rs string) (Response, error) {
	// array
	var arrobjgetPasienMapIcd []GetPasienMapIcd
	// object
	var objgetPasienMapIcd GetPasienMapIcd

	var res Response

	con := database.CreateCon()

	//getPasienMapIcd
	getPasienMapIcd := "select desano, nama_desa, id_rs, waktu from pasien_maps where id_provinsi=? and id_rs=? order by desano asc"
	rowsgetPasienMapIcd, err := con.Query(getPasienMapIcd, id_provinsi, id_rs)
	if err != nil {
		fmt.Println("Data getPasienMapIcd has been successfully loaded.")
		return res, err
	}

	defer rowsgetPasienMapIcd.Close()
	//close

	for rowsgetPasienMapIcd.Next() {
		err = rowsgetPasienMapIcd.Scan(&objgetPasienMapIcd.Desano,&objgetPasienMapIcd.Nama_desa, &objgetPasienMapIcd.Id_rs,&objgetPasienMapIcd.Waktu)
		if err != nil {
			fmt.Println("Data getPasienMapIcd has been successfully loaded on Rows Next.", err)
			return res, err
		}

		arrobjgetPasienMapIcd = append(arrobjgetPasienMapIcd, objgetPasienMapIcd)
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = arrobjgetPasienMapIcd
	
	return res, nil
}

func GetPasienMapIcdDataV2(id_provinsi, id_rs string) ([]GetPasienMapIcd, error) {
	// array
	var arrobjgetPasienMapIcd []GetPasienMapIcd
	// object
	var objgetPasienMapIcd GetPasienMapIcd

	con := database.CreateCon()

	//getPasienMapIcd
	getPasienMapIcd := "select desano, nama_desa, id_rs, waktu from pasien_maps where id_provinsi=? and id_rs=? order by desano asc"
	rowsgetPasienMapIcd, err := con.Query(getPasienMapIcd, id_provinsi, id_rs)
	if err != nil {
		fmt.Println("Data getPasienMapIcd has been successfully loaded.")
		return arrobjgetPasienMapIcd, err
	}

	defer rowsgetPasienMapIcd.Close()
	//close

	for rowsgetPasienMapIcd.Next() {
		err = rowsgetPasienMapIcd.Scan(&objgetPasienMapIcd.Desano,&objgetPasienMapIcd.Nama_desa, &objgetPasienMapIcd.Id_rs,&objgetPasienMapIcd.Waktu)
		if err != nil {
			fmt.Println("Data getPasienMapIcd has been successfully loaded on Rows Next.", err)
			return arrobjgetPasienMapIcd, err
		}

		arrobjgetPasienMapIcd = append(arrobjgetPasienMapIcd, objgetPasienMapIcd)
	}
	
	return arrobjgetPasienMapIcd, nil
}