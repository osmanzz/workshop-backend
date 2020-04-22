package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/osmanzz/workshop-backend/model"
	"github.com/osmanzz/workshop-backend/pkgs"
)

func InsertMahasiswa(writer http.ResponseWriter, request *http.Request) {
	req := &model.UserDataRequest{}
	response := &model.Response{}
	err := json.NewDecoder(request.Body).Decode(req)
	if err != nil {
		log.Printf("err:%+v\n", err)
	}
	err = pkgs.InsertUserData(req)
	if err != nil {
		log.Printf("err:%+v\n", err)
	}

	response.Message = "Yeay Berhasil."
	byteData, _ := json.Marshal(response)
	writer.Header().Add("Content-type", "application/json")
	writer.Write(byteData)
}
