package model

type UserData struct {
	ID       int64  `db:"id"json:"id"`
	Name     string `db:"name"json:"name"`
	Password string `db:"password"json:"password"`
	Alamat   string `db:"alamat"json:"alamat"`
}

type PerpustakaanData struct {
	ID      int64      `db:"id"json:"id"`
	MhsID   int64      `db:"mhs_id"json:"mhs_id"`
	Buku    string     `db:"buku"json:"buku"`
	Tanggal string	   `db:"tanggal"json:"tanggal"`
}

type ResponseData struct {
	UserData []*User `json:"user_data"`
}
type User struct {
	ID           int64           `db:"id"json:"id"`
	Name         string          `db:"name"json:"name"`
	Password     string          `db:"password"json:"password"`
	Alamat       string          `db:"alamat"json:"alamat"`
	Perpustakaan []*Perpustakaan `json:"perpustakaan"`
}

type Perpustakaan struct {
	ID      int64      `db:"id"json:"id"`
	MhsID   int64      `db:"mhs_id"json:"mhs_id"`
	Buku    string     `db:"buku"json:"buku"`
	Tanggal string     `db:"tanggal"json:"tanggal"`
}

type UserDataRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Alamat   string `json:"alamat"`
}

type Response struct {
	Message string `json:"message"`
}
