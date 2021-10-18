package models

import (
	// "encoding/json"
	"fmt"
	"golang-rest-api/database"
	// "io/ioutil"
	// "log"
	"net/http"
)

type DataPariwisata struct {
	PariwisataId      int    `json:"pariwisata_id"`
	PariwisataNama  string `json:"pariwisata_nama"`
	PariwisataLokasi string `json:"pariwisata_lokasi"`
	PariwisataKeterangan string `json:"pariwisata_keterangan"`
	Populasi string `json:"populasi"`
}

func PariwisataData() (Response, error) {
	var obj DataPariwisata
	var arrobj []DataPariwisata
	var res Response

	con := database.CreateCon()

	sqlStatement := "SELECT * FROM pariwisata order by id desc"

	rows, err := con.Query(sqlStatement)


	if err != nil {
		fmt.Println("masuk sini")
		return res, err
	}
	
	//defer ini bisa jadi harus di bawah return res
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.PariwisataId, &obj.PariwisataNama, &obj.PariwisataLokasi, &obj.PariwisataKeterangan, &obj.Populasi)

		if err != nil {
			fmt.Println("masuk sini 2 ", err)
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = arrobj
	fmt.Println("berhasil 12 ", res.Status)
	return res, nil
}