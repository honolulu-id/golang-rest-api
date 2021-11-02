package models

import (
	"fmt"
	"golang-rest-api/database"
	"net/http"
)

type getGroupByDesano struct {
	Desano string `json:"desano"`
}

type getLatLongDesano struct {
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}


type getPasienMapIcd struct {
	Desano string `json:"desano"`
	Nama_desa string `json:"nama_desa"`
	Id_rs string `json:"id_rs"`
	Waktu string `json:"waktu"`
}

type getPenyakitByKelurahan struct {
	Nama_penyakit string `json:"nama_penyakit"`
	Kode string `json:"kode"`
	Kode_icd string `json:"kode_icd"`
	Id_rs string `json:"id_rs"`
	Kode_kelurahan string `json:"kode_kelurahan"`
	Waktu string `json:"waktu"`
}

func GetGroupByDesanoData() (Response, error) {
	// array
	var arrobjgetGroupByDesano []getGroupByDesano

	// object
	var objgetGroupByDesano getGroupByDesano

	var res Response

	con := database.CreateCon()

	//getGroupByDesano
	getGroupByDesano := "select desano from kelurahan_maps where id_provinsi=11 and id_rs=1 group by desano desc"
	rowsgetGroupByDesano, err := con.Query(getGroupByDesano)
	if err != nil {
		fmt.Println("Data getGroupByDesano has been successfully loaded.")
		return res, err
	}

	defer rowsgetGroupByDesano.Close()
	//close

	for rowsgetGroupByDesano.Next() {
		err = rowsgetGroupByDesano.Scan(&objgetGroupByDesano.Desano)
		if err != nil {
			fmt.Println("Data getGroupByDesano has been successfully loaded on Rows Next.", err)
			return res, err
		}

		arrobjgetGroupByDesano = append(arrobjgetGroupByDesano, objgetGroupByDesano)
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = arrobjgetGroupByDesano
	fmt.Println("Success Load Data ", res.Status)
	return res, nil
}

func GetPasienMapIcdData() (Response, error) {
	// array
	var arrobjgetPasienMapIcd []getPasienMapIcd
	// object
	var objgetPasienMapIcd getPasienMapIcd

	var res Response

	con := database.CreateCon()

	//getPasienMapIcd
	getPasienMapIcd := "select desano, nama_desa, id_rs, waktu from pasien_maps where id_provinsi=11 and id_rs=1 order by desano asc"
	rowsgetPasienMapIcd, err := con.Query(getPasienMapIcd)
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
	fmt.Println("Success Load Data ", res.Status)
	return res, nil
}

func GetPenyakitByKelurahanData() (Response, error) {
	// array
	var arrobjgetPenyakitByKelurahan []getPenyakitByKelurahan

	// object
	var objgetPenyakitByKelurahan getPenyakitByKelurahan

	var res Response

	con := database.CreateCon()

	//getPenyakitByKelurahan
	getPenyakitByKelurahan := "select nama_penyakit, kode, kode_icd, id_rs, kode_kelurahan,waktu from penyakit_maps where kode_icd=11134 and id_rs=1 and id_provinsi=11"
	rowsgetPenyakitByKelurahan, err := con.Query(getPenyakitByKelurahan)
	if err != nil {
		fmt.Println("Data getPenyakitByKelurahan has been successfully loaded.")
		return res, err
	}

	defer rowsgetPenyakitByKelurahan.Close()
	//close

	for rowsgetPenyakitByKelurahan.Next() {
		err = rowsgetPenyakitByKelurahan.Scan(&objgetPenyakitByKelurahan.Nama_penyakit,&objgetPenyakitByKelurahan.Kode, &objgetPenyakitByKelurahan.Kode_icd,&objgetPenyakitByKelurahan.Id_rs,&objgetPenyakitByKelurahan.Kode_kelurahan,&objgetPenyakitByKelurahan.Waktu)
		if err != nil {
			fmt.Println("Data getPenyakitByKelurahan has been successfully loaded on Rows Next.", err)
			return res, err
		}

		arrobjgetPenyakitByKelurahan = append(arrobjgetPenyakitByKelurahan, objgetPenyakitByKelurahan)
	}


	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = arrobjgetPenyakitByKelurahan
	fmt.Println("Success Load Data ", res.Status)
	return res, nil
}

func GetLatLongDesanoData() (Response, error) {
	// array
	var arrobjgetLatLongDesano []getLatLongDesano

	// object
	var objgetLatLongDesano getLatLongDesano

	var res Response

	con := database.CreateCon()

	//getGroupByDesano
	getLatLongDesano := "select latitude, longitude from kelurahan_maps where id_provinsi=11 and id_rs=1"
	rowsgetLatLongDesano, err := con.Query(getLatLongDesano)
	if err != nil {
		fmt.Println("Data getLatLongDesano has been successfully loaded.")
		return res, err
	}

	defer rowsgetLatLongDesano.Close()
	//close

	for rowsgetLatLongDesano.Next() {
		err = rowsgetLatLongDesano.Scan(&objgetLatLongDesano.Latitude,&objgetLatLongDesano.Longitude)
		if err != nil {
			fmt.Println("Data getLatLongDesano has been successfully loaded on Rows Next.", err)
			return res, err
		}

		arrobjgetLatLongDesano = append(arrobjgetLatLongDesano, objgetLatLongDesano)
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = arrobjgetLatLongDesano
	fmt.Println("Success Load Data ", res.Status)
	return res, nil
}