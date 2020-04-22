package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/workshop/pkgs"
)

func GetAllData(writer http.ResponseWriter, request *http.Request) {
	resp, err := pkgs.GetAllUserData()
	if err != nil {
		log.Println("err: %+v", err)
	}

	byteData, _ := json.Marshal(resp)
	writer.Header().Add("Content-type", "application/json")
	writer.Write(byteData)
}
