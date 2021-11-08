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

func GetGroupByDesanoData(id_provinsi, id_rs string) (Response, error) {
	// array
	var arrobjgetGroupByDesano []GetGroupByDesano

	// object
	var objgetGroupByDesano GetGroupByDesano

	var res Response

	con := database.CreateCon()

	//getGroupByDesano
	getGroupByDesano := "select desano from kelurahan_maps where id_provinsi=? and id_rs=? group by desano desc"
	rowsgetGroupByDesano, err := con.Query(getGroupByDesano, id_provinsi, id_rs)
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
	// fmt.Println("Success Load Data ", res.Status)
	return res, nil
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

func GetPenyakitByKelurahanData(kode_icd, kode_kelurahan, id_rs string) (Response, error) {
	// array
	var arrobjgetPenyakitByKelurahan []GetPenyakitByKelurahan

	// object
	var objgetPenyakitByKelurahan GetPenyakitByKelurahan

	var res Response

	con := database.CreateCon()

	//getPenyakitByKelurahan
	getPenyakitByKelurahan := `select 
								nama_penyakit, kode, kode_icd, id_rs, kode_kelurahan,jml_pasien 
							   from 
							   	penyakit_maps where kode_icd=? 
							   and 
							   	kode_kelurahan=? and id_rs=? order by id desc limit 1`
	rowsgetPenyakitByKelurahan, err := con.Query(getPenyakitByKelurahan, kode_icd, kode_kelurahan, id_rs)

	if err != nil {
		fmt.Println("Data getPenyakitByKelurahan Failed loaded.")
		return res, err
	}

	defer rowsgetPenyakitByKelurahan.Close()
	//close

	// fmt.Println("id rs", id_rs, " provinsi", id_provinsi, "kode icd", kode_icd)

	for rowsgetPenyakitByKelurahan.Next() {
		err = rowsgetPenyakitByKelurahan.Scan(&objgetPenyakitByKelurahan.Nama_penyakit,&objgetPenyakitByKelurahan.Kode, &objgetPenyakitByKelurahan.Kode_icd,&objgetPenyakitByKelurahan.Id_rs,&objgetPenyakitByKelurahan.Kode_kelurahan,&objgetPenyakitByKelurahan.Jumlah_Pasien)
		if err != nil {
			fmt.Println("Data getPenyakitByKelurahan has been successfully loaded on Rows Next.", err)
			return res, err
		}

		arrobjgetPenyakitByKelurahan = append(arrobjgetPenyakitByKelurahan, objgetPenyakitByKelurahan)
	}


	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = arrobjgetPenyakitByKelurahan
	// fmt.Println("Success Load Data ", kode_icd)
	return res, nil
}

func GetPenyakitByKelurahanDataV2(kode_icd, kode_kelurahan, id_rs string) (GetPenyakitByKelurahan, error) {

	// object
	var objgetPenyakitByKelurahan GetPenyakitByKelurahan

	con := database.CreateCon()

	//getPenyakitByKelurahan
	getPenyakitByKelurahan := `select 
								nama_penyakit, kode, kode_icd, id_rs, kode_kelurahan,jml_pasien 
							   from 
							   	penyakit_maps where kode_icd=? 
							   and 
							   	kode_kelurahan=? and id_rs=? order by id asc`
	rowsgetPenyakitByKelurahan, err := con.Query(getPenyakitByKelurahan, kode_icd, kode_kelurahan, id_rs)

	if err != nil {
		fmt.Println("Data getPenyakitByKelurahan Failed loaded.", err)
		return objgetPenyakitByKelurahan, err
	}

	defer rowsgetPenyakitByKelurahan.Close()
	//close
	
	err = rowsgetPenyakitByKelurahan.Scan(&objgetPenyakitByKelurahan.Nama_penyakit,&objgetPenyakitByKelurahan.Kode, &objgetPenyakitByKelurahan.Kode_icd,&objgetPenyakitByKelurahan.Id_rs,&objgetPenyakitByKelurahan.Kode_kelurahan,&objgetPenyakitByKelurahan.Jumlah_Pasien)
	if err != nil {
		return objgetPenyakitByKelurahan, err
	}
	
	return objgetPenyakitByKelurahan, err
}

func GetLatLongDesanoData(desano, id_rs string) (Response, error) {
	// array
	var arrobjgetLatLongDesano []GetLatLongDesano

	// object
	var objgetLatLongDesano GetLatLongDesano

	var res Response

	con := database.CreateCon()

	//getGroupByDesano
	getLatLongDesano := "select latitude, longitude from kelurahan_maps where desano= ? and id_rs= ?"
	rowsgetLatLongDesano, err := con.Query(getLatLongDesano, desano, id_rs)
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
	
	return res, nil
}