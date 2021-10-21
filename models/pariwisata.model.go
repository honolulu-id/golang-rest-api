package models

import (
	"fmt"
	"golang-rest-api/database"
	"net/http"
)

type DataPariwisata struct {
	PariwisataId      string    `json:"pariwisata_id"`
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
	defer database.CloseDatabase()

	//run sql di sini
	sqlStatement := "SELECT * FROM pariwisata order by pariwisata_id desc"
	rows, err := con.Query(sqlStatement)
	if err != nil {
		fmt.Println("Data has been successfully loaded.")
		return res, err
	}

	defer rows.Close()
	//close

	for rows.Next() {
		err = rows.Scan(&obj.PariwisataId, &obj.PariwisataNama, &obj.PariwisataLokasi, &obj.PariwisataKeterangan, &obj.Populasi)
		// err = rows.Scan(&arrobj)
		if err != nil {
			fmt.Println("Data has been successfully loaded on Rows Next.", err)
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = arrobj
	fmt.Println("Success Load Data ", res.Status)
	return res, nil
}