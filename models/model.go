package models

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	//"time"
)

var (
	Intelligencedatas []Intelligencedata
)

type Intelligencedata struct {
	Gid         int
	Domain      string
	Path        string
	Discreption string
	Flag        int
}

var (
	Users map[string]*User
)

type User struct {
	Uid         int
	Username    string
	Passwd      string
	Discreption string
	Pid         int
}

var (
	Permissions map[string]*Permission
)

type Permission struct {
	Pid         int
	Discreption string
}

func Init() []Intelligencedata {

	//Intelligencedatas = make(map[string]*Intelligencedata)

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/goweb?charset=utf8")
	checkErr(err)
	var Intelligencedata Intelligencedata
	var Discreption []byte

	//查询数据
	rows, err := db.Query("SELECT * FROM intelligencedata")
	checkErr(err)
	for rows.Next() {

		err = rows.Scan(&Intelligencedata.Gid, &Intelligencedata.Domain, &Intelligencedata.Path, &Discreption, &Intelligencedata.Flag)
		checkErr(err)
		Intelligencedata.Discreption = string(Discreption)
		Intelligencedatas = append(Intelligencedatas, Intelligencedata)
	}

	db.Close()
	return Intelligencedatas
}

func Search(domain string) []Intelligencedata {

	var Datas []Intelligencedata

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/goweb?charset=utf8")
	checkErr(err)
	var Intelligencedata Intelligencedata
	var Discreption []byte

	sqlsten := "SELECT * FROM intelligencedata where domain like '%" + domain + "%'"
	fmt.Println(sqlsten)
	fmt.Println(domain)
	//查询数据
	rows, err := db.Query(sqlsten)
	checkErr(err)
	for rows.Next() {

		err = rows.Scan(&Intelligencedata.Gid, &Intelligencedata.Domain, &Intelligencedata.Path, &Discreption, &Intelligencedata.Flag)
		checkErr(err)
		Intelligencedata.Discreption = string(Discreption)
		Datas = append(Datas, Intelligencedata)
	}
	db.Close()
	fmt.Println(Datas)
	return Datas
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
