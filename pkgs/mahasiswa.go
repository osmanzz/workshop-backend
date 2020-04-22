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

	err = db.Select(&res, "select * from perpustakaan where mhs_id=?", id)
	if err != nil {
		return res, err
	}

	return res, nil
}

func InsertUserData(data *model.UserDataRequest) error {
	db, err := InitDB()
	if err != nil {
		return err
	}
	tx := db.MustBegin()
	tx.MustExec("insert into mahasiswa(name,password,alamat) values (?,?,?)", data.Name, data.Password, data.Alamat)
	tx.Commit()
	// err = db.Get(&userID,"insert into mahasiswa(name,password,alamat) values ($1,$2,$3) returning id", data.Name, data.Password, data.Alamat )
	// if err != nil {
	// 	return err
	// }
	return nil
}
