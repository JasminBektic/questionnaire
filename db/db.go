package db

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
)

func ReadConfig() map[string]string {
	read, err := ioutil.ReadFile("config/db.json")
	if err != nil {
		panic(err)
	}

	var config map[string]string

	err = json.Unmarshal(read, &config)
	if err != nil {
		panic(err)
	}

	return config
}

func Open() (*sql.DB, error) {
	c := ReadConfig()

	db, err := sql.Open("mysql", c["db_user"]+":"+c["db_password"]+"@tcp("+c["db_host"]+":"+c["db_port"]+")/"+c["db_name"])
	if err != nil {
		panic(err)
	}

	return db, err
}
