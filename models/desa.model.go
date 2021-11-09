package models

import (
	"fmt"
	"golang-rest-api/database"
	"net/http"
)

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

func GetLatLongDesanoDataV2(desano, id_rs string) ([]GetLatLongDesano, error) {
	// array
	var arrobjgetLatLongDesano []GetLatLongDesano

	// object
	var objgetLatLongDesano GetLatLongDesano

	con := database.CreateCon()

	//getGroupByDesano
	getLatLongDesano := "select latitude, longitude from kelurahan_maps where desano= ? and id_rs= ?"
	rowsgetLatLongDesano, err := con.Query(getLatLongDesano, desano, id_rs)
	if err != nil {
		fmt.Println("Data getLatLongDesano has been successfully loaded.")
		return arrobjgetLatLongDesano, err
	}

	defer rowsgetLatLongDesano.Close()
	//close

	for rowsgetLatLongDesano.Next() {
		err = rowsgetLatLongDesano.Scan(&objgetLatLongDesano.Latitude,&objgetLatLongDesano.Longitude)
		if err != nil {
			fmt.Println("Data getLatLongDesano has been successfully loaded on Rows Next.", err)
			return arrobjgetLatLongDesano, err
		}

		arrobjgetLatLongDesano = append(arrobjgetLatLongDesano, objgetLatLongDesano)
	}
	
	return arrobjgetLatLongDesano, nil
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