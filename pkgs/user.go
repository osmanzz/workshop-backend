package pkgs

import "github.com/workshop/model"

func GetAllUserData() ([]*model.UserData, error) {
	res := make([]*model.UserData, 0)
	db, err := InitDB()
	if err != nil {
		return res, err
	}
	err = db.Select(&res, "select * from ws_user")
	if err != nil {
		return res, err
	}

	return res, nil
}
