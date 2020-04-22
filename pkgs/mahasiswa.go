package pkgs

import (
	"github.com/osmanzz/workshop-backend/model"
)

func GetAllUserData() ([]*model.UserData, error) {
	res := make([]*model.UserData, 0)
	db, err := InitDB()
	if err != nil {
		return res, err
	}
	err = db.Select(&res, "select * from mahasiswa")
	if err != nil {
		return res, err
	}

	return res, nil
}

func GetAllPerpustakaanDataByUserID(id int64) ([]*model.PerpustakaanData, error) {
	res := make([]*model.PerpustakaanData, 0)
	db, err := InitDB()
	if err != nil {
		return res, err
	}
	err = db.Select(&res, "select * from perpustakaan where mhs_id=$1", id)
	if err != nil {
		return res, err
	}

	return res, nil
}
