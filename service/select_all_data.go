package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/osmanzz/workshop-backend/model"
	"github.com/osmanzz/workshop-backend/pkgs"
)

func GetAllData(writer http.ResponseWriter, request *http.Request) {
	result := &model.ResponseData{}
	resp, err := pkgs.GetAllUserData()
	if err != nil {
		log.Println("err: %+v", err)
	}

	for _, data := range resp {
		temp := &model.User{
			ID:       data.ID,
			Name:     data.Name,
			Password: data.Password,
			Alamat:   data.Alamat,
		}

		perpustakaan, err := pkgs.GetAllPerpustakaanDataByUserID(data.ID)
		if err != nil {
			log.Println("err: %+v", err)
		}
		for _, dt := range perpustakaan {
			tp := &model.Perpustakaan{
				ID:      dt.ID,
				MhsID:   dt.MhsID,
				Buku:    dt.Buku,
				Tanggal: dt.Tanggal,
			}
			temp.Perpustakaan = append(temp.Perpustakaan, tp)
		}
		result.UserData = append(result.UserData, temp)

	}
	byteData, _ := json.Marshal(result)
	writer.Header().Add("Content-type", "application/json")
	writer.Write(byteData)
}
