package models

import (
	"fmt"
	"golang-rest-api/database"
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
		fmt.Println("Data has been successfully loaded.")
		return res, err
	}

	//defer ini bisa jadi harus di bawah return res
	//setiap koneksi yang mengakses kuery harus di close setelah tidak ada transaksi 
	//sehingga tidak memberatkan open connection selanjutnya
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.PariwisataId, &obj.PariwisataNama, &obj.PariwisataLokasi, &obj.PariwisataKeterangan, &obj.Populasi)

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