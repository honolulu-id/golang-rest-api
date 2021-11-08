package controllers

import (
	"fmt"
	echo "github.com/labstack/echo/v4"
	"golang-rest-api/models"
	"net/http"
	// "encoding/json"
	// "reflect"
)

type GroupByDesanoData struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []struct {
		Desano string `json:"desano"`
	} `json:"data"`
}

type NilaiSama struct {
	Desano string `json:"desano"`
    NamaDesa string `json:"nama_desa"`
    IDRs string `json:"id_rs"`
    // Waktu string json:"waktu"
}

type PenyakitFormatvOne struct {
	NamaPenyakit string `json:"nama_penyakit"`
	KoDe string `json:"kode"`
	KodeIcd string `json:"kode_icd"`
}

type PenyakitFormatvTwo struct {
	JmlPasien string `json:"jml_pasien"`
}

type Coordiantes struct {
	Longitude string `json:"longitude"`
	Latitude string `json:"latitude"`
}

func HeatMapDataList(c echo.Context) error {

	idProvinsi := c.QueryParams().Get("id_provinsi")
	idRs := c.QueryParams().Get("id_rs")
	Kode_Icd := c.QueryParams().Get("kode_icd")

	resultGetGroupByDesanoData, err := models.GetGroupByDesanoData(idProvinsi, idRs)
	resultGetPasienMapIcdData, err := models.GetPasienMapIcdData(idProvinsi, idRs)

	var nilai_sama []NilaiSama
	var obj NilaiSama
	var Coordiante []Coordiantes
	var objCoordiantes Coordiantes
	var PenyakitFormatv_1 []PenyakitFormatvOne
	var objPenyakitFormatv_1 PenyakitFormatvOne
	var PenyakitFormatv_2 []PenyakitFormatvTwo
	var objPenyakitFormatv_2 PenyakitFormatvTwo

	for _, MapIcdData := range resultGetPasienMapIcdData.Data.([]models.GetPasienMapIcd) {
		var Kode_Fix = MapIcdData.Desano
		for _, DesanoData := range resultGetGroupByDesanoData.Data.([]models.GetGroupByDesano) {
			if DesanoData.Desano == Kode_Fix {
				obj.Desano = MapIcdData.Desano
				obj.NamaDesa = MapIcdData.Nama_desa
				obj.IDRs = MapIcdData.Id_rs
				nilai_sama = append(nilai_sama, obj)
			}
		}
		// fmt.Printf("DesanoData %v\n", nilai_sama)
	}

	for _, LoopNilaiSama := range nilai_sama {
	var Id_Desano = LoopNilaiSama.Desano
	var Id_RumahSakit = LoopNilaiSama.IDRs

	Kelurahan, _ := models.GetLatLongDesanoData(Id_Desano,idRs)
	Penyakit, _ := models.GetPenyakitByKelurahanData(Kode_Icd, Id_Desano, Id_RumahSakit)

		for _, Formatv_1 := range Penyakit.Data.([]models.GetPenyakitByKelurahan) {

			objPenyakitFormatv_1.NamaPenyakit = Formatv_1.Nama_penyakit
			objPenyakitFormatv_1.KoDe = Formatv_1.Kode
			objPenyakitFormatv_1.KodeIcd = Formatv_1.Kode_icd
			PenyakitFormatv_1 = append(PenyakitFormatv_1, objPenyakitFormatv_1)
			// fmt.Printf("Penyakit %v\n", objPenyakitFormatv_1)
		}

		for _, Formatv_2 := range Penyakit.Data.([]models.GetPenyakitByKelurahan) {
			objPenyakitFormatv_2.JmlPasien = Formatv_2.Jumlah_Pasien
			PenyakitFormatv_2 = append(PenyakitFormatv_2, objPenyakitFormatv_2)
			// fmt.Printf("Id_RumahSakit %v\n", PenyakitFormatv_2)
		}

		for _, Koordinat := range Kelurahan.Data.([]models.GetLatLongDesano) {
			objCoordiantes.Longitude = Koordinat.Longitude
			objCoordiantes.Latitude = Koordinat.Latitude
			Coordiante = append(Coordiante, objCoordiantes)
			fmt.Printf("Coordiantes %v\n", Koordinat)
		}

	// reformat penyakit here
	// fmt.Printf("Id_RumahSakit %v\n", Id_RumahSakit)

	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, Coordiante)

}
