package models

import (
	"fmt"
	"golang-rest-api/database"
	"net/http"
)

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
								nama_penyakit, 
								kode, 
								kode_icd, 
								id_rs, 
								kode_kelurahan,
								COALESCE(jml_pasien, 0) as jml_pasien
							   from 
							   	penyakit_maps 
							   where 
							   	kode_icd=? 
							   and 
							   	kode_kelurahan=? 
							   and 
							   	id_rs=? order by id asc`
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