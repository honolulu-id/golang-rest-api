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

func HeatMapDataList(c echo.Context) error {

	idProvinsi := c.QueryParams().Get("id_provinsi")
	idRs := c.QueryParams().Get("id_rs")
	// KodeIcd := c.QueryParams().Get("kode_icd")

	resultGetGroupByDesanoData, err := models.GetGroupByDesanoData(idProvinsi, idRs)
	resultGetPasienMapIcdData, err := models.GetPasienMapIcdData(idProvinsi, idRs)

	//array coordiantes [] type data string default kosong{}
	// var coordinates []string
	// nilai_sama := make([]string, 0)
	var nilai_sama []NilaiSama
	var obj NilaiSama

	for _, MapIcdData := range resultGetPasienMapIcdData.Data.([]models.GetPasienMapIcd) {
		var Kode_Fix = MapIcdData.Desano
		for _, DesanoData := range resultGetGroupByDesanoData.Data.([]models.GetGroupByDesano) {
			if DesanoData.Desano == Kode_Fix {
				obj.Desano = MapIcdData.Desano
				obj.NamaDesa = MapIcdData.Nama_desa
				obj.IDRs = MapIcdData.Id_rs
				// nilai_sama = append(nilai_sama, MapIcdData.Desano, MapIcdData.Nama_desa, MapIcdData.Id_rs)
				nilai_sama = append(nilai_sama, obj)
			}
			// fmt.Printf("Kode_Fix %v\n", Kode_Fix)
		}
		fmt.Printf("DesanoData %v\n", MapIcdData)
	}

	// for _, LoopNilaiSama := range nilai_sama {
	// resultGetLatLongDesanoData, err := models.GetLatLongDesanoData(LoopNilaiSama,idRs)
	// resultGetPenyakitByKelurahanData, err := models.GetPenyakitByKelurahanData(KodeIcd, idProvinsi, idRs)
	// fmt.Printf("LoopNilaiSama %v\n", nilai_sama)
	// fmt.Printf("LoopNilaiSama %v\n", resultGetLatLongDesanoData)
	// }

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, resultGetPasienMapIcdData)

}
