package controllers

import (
	"golang-rest-api/models"
	"net/http"
	"time"

	echo "github.com/labstack/echo/v4"
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

type Properties struct {
	Desano string `json:"desano"`
    NamaDesa string `json:"nama_desa"`
    IDRs string `json:"id_rs"`
    JmlPasien string `json:"jml_pasien"`
}

type Coordiantes struct {
	Longitude string `json:"longitude"`
	Latitude string `json:"latitude"`
}

type Features struct {
	Type string `json:"type"`
	Penyakit interface{} `json:"penyakit"`
	Properties interface{} `json:"properties"`
	Geometry interface{} `json:"geometry"` 
}

type Geometry struct {
	Type string `json:"type"`
	Coordinates interface{} `json:"coordinates"`
}

func HeatMapDataList(c echo.Context) error {
	
	idProvinsi := c.QueryParams().Get("id_provinsi")
	idRs := c.QueryParams().Get("id_rs")
	Kode_Icd := c.QueryParams().Get("kode_icd")

	resultGetGroupByDesanoData, _ := models.GetGroupByDesanoData(idProvinsi, idRs)
	resultGetPasienMapIcdData, _ := models.GetPasienMapIcdDataV2(idProvinsi, idRs)

	var nilai_sama []NilaiSama
	var obj NilaiSama
	var Coordiante []Coordiantes
	var objCoordiantes Coordiantes
	var objPenyakitFormatv_1 PenyakitFormatvOne
	var featuresObj Features
	var featuresArrObj []Features 
	var geometry Geometry

	var propertiesObj Properties
	
	for _, MapIcdData := range resultGetPasienMapIcdData {
		var Kode_Fix = MapIcdData.Desano

		for _, DesanoData := range resultGetGroupByDesanoData.Data.([]models.GetGroupByDesano) {
			if DesanoData.Desano == Kode_Fix {
				obj.Desano = MapIcdData.Desano
				obj.NamaDesa = MapIcdData.Nama_desa
				obj.IDRs = MapIcdData.Id_rs
				nilai_sama = append(nilai_sama, obj)
			}
		}

	}

	for _, loopnilai := range nilai_sama {

		var Id_Desano = loopnilai.Desano
		var Id_RumahSakit = loopnilai.IDRs
		Kelurahan, _ := models.GetLatLongDesanoDataV2(Id_Desano,idRs)

		for _, Koordinat := range Kelurahan {
			objCoordiantes.Longitude = Koordinat.Longitude
			objCoordiantes.Latitude = Koordinat.Latitude
			Coordiante = append(Coordiante, objCoordiantes)
		}

		Penyakit, _ := models.GetPenyakitByKelurahanDataV2(Kode_Icd, Id_Desano, Id_RumahSakit)
		objPenyakitFormatv_1.NamaPenyakit = Penyakit.Nama_penyakit
		objPenyakitFormatv_1.KoDe = Penyakit.Kode
		objPenyakitFormatv_1.KodeIcd = Penyakit.Kode_icd

		propertiesObj.Desano = Id_Desano
		propertiesObj.IDRs = Id_RumahSakit
		propertiesObj.NamaDesa = loopnilai.NamaDesa
		propertiesObj.JmlPasien = Penyakit.Jumlah_Pasien

		geometry.Type = "MultiPolygon"
		geometry.Coordinates = Coordiante

		featuresObj.Type = "Feature"
		featuresObj.Penyakit = objPenyakitFormatv_1
		featuresObj.Properties = propertiesObj
		featuresObj.Geometry = geometry

		featuresArrObj = append(featuresArrObj, featuresObj)
	}

	var res models.ResponseApi

	if nilai_sama == nil {
		res.Status = false
		res.Message = "gagal mendapatkan data"
		res.Data = nilai_sama
		return c.JSON(http.StatusOK, res)
	}

	res.Status = true
	res.Message = "Berhasil mendapatkan data"
	res.Data = featuresArrObj

	time.Sleep(100 * time.Second)
	return c.JSON(http.StatusOK, res)
}

func HeatMapDataListBackup(c echo.Context) error {
	
	idProvinsi := c.QueryParams().Get("id_provinsi")
	idRs := c.QueryParams().Get("id_rs")
	Kode_Icd := c.QueryParams().Get("kode_icd")

	resultGetGroupByDesanoData, err := models.GetGroupByDesanoData(idProvinsi, idRs)
	resultGetPasienMapIcdData, err := models.GetPasienMapIcdData(idProvinsi, idRs)

	var nilai_sama []NilaiSama
	var obj NilaiSama
	var Coordiante []Coordiantes
	var objCoordiantes Coordiantes
	// var PenyakitFormatv_1 []PenyakitFormatvOne
	var objPenyakitFormatv_1 PenyakitFormatvOne
	// var PenyakitFormatv_2 []PenyakitFormatvTwo
	// var objPenyakitFormatv_2 PenyakitFormatvTwo
	
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

	var res models.ResponseApi

	if nilai_sama == nil {
		res.Status = false
		res.Message = "gagal mendapatkan data"
		res.Data = nilai_sama
		return c.JSON(http.StatusOK, res)
	}

	var featuresObj Features
	var featuresArrObj []Features 
	var geometry Geometry

	var propertiesObj Properties
	// var propertiesArrObj []Properties
	
	for _, LoopNilaiSama := range nilai_sama {
		var Id_Desano = LoopNilaiSama.Desano
		var Id_RumahSakit = LoopNilaiSama.IDRs

		// fmt.Println("id desano", Id_Desano)

		// Kelurahan, _ := models.GetLatLongDesanoData(Id_Desano,idRs)
		Kelurahan, _ := models.GetLatLongDesanoDataV2(Id_Desano,idRs)
		// Penyakit, _ := models.GetPenyakitByKelurahanData(Kode_Icd, Id_Desano, Id_RumahSakit)
		Penyakit, _ := models.GetPenyakitByKelurahanDataV2(Kode_Icd, Id_Desano, Id_RumahSakit)
		// fmt.Println("data penyakit", Penyakit)

		for _, Koordinat := range Kelurahan {
			objCoordiantes.Longitude = Koordinat.Longitude
			objCoordiantes.Latitude = Koordinat.Latitude
			Coordiante = append(Coordiante, objCoordiantes)
		}

		// for _, Formatv_1 := range Penyakit.Data.([]models.GetPenyakitByKelurahan) {
		// 	// fmt.Printf("Penyakit %v\n", objPenyakitFormatv_1)
		// 	if Formatv_1.Nama_penyakit != "" {
		// 		objPenyakitFormatv_1.NamaPenyakit = Formatv_1.Nama_penyakit
		// 		objPenyakitFormatv_1.KoDe = Formatv_1.Kode
		// 		objPenyakitFormatv_1.KodeIcd = Formatv_1.Kode_icd
		// 		PenyakitFormatv_1 = append(PenyakitFormatv_1, objPenyakitFormatv_1)
		// 	}
		// }

		objPenyakitFormatv_1.NamaPenyakit = Penyakit.Nama_penyakit
		objPenyakitFormatv_1.KoDe = Penyakit.Kode
		objPenyakitFormatv_1.KodeIcd = Penyakit.Kode_icd

		// for _, Formatv_2 := range Penyakit.Data.([]models.GetPenyakitByKelurahan) {
		// 	objPenyakitFormatv_2.JmlPasien = Formatv_2.Jumlah_Pasien
		// 	PenyakitFormatv_2 = append(PenyakitFormatv_2, objPenyakitFormatv_2)

		// 	propertiesObj.Desano = Id_Desano
		// 	propertiesObj.IDRs = Id_RumahSakit
		// 	propertiesObj.NamaDesa = LoopNilaiSama.NamaDesa
		// 	propertiesObj.JmlPasien = Formatv_2.Jumlah_Pasien
		// 	propertiesArrObj = append(propertiesArrObj, propertiesObj)
		// }

		propertiesObj.Desano = Id_Desano
		propertiesObj.IDRs = Id_RumahSakit
		propertiesObj.NamaDesa = LoopNilaiSama.NamaDesa
		propertiesObj.JmlPasien = Penyakit.Jumlah_Pasien

		geometry.Type = "MultiPolygon"
		geometry.Coordinates = Coordiante

		featuresObj.Type = "Feature"
		featuresObj.Penyakit = objPenyakitFormatv_1
		featuresObj.Properties = propertiesObj
		featuresObj.Geometry = geometry

		featuresArrObj = append(featuresArrObj, featuresObj)

	}

	time.Sleep(10 * time.Second)
	return c.JSON(http.StatusOK, Coordiante)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	res.Status = true
	res.Message = "Berhasil mendapatkan data"
	res.Data = featuresArrObj
	return c.JSON(http.StatusOK, res)

}
