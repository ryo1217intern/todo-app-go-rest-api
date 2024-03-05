package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
)

//NewDBとCloseDBを使う処理を行い。作成したデータベースにmodelで作成したスキーマを反映させる.
func main() {
	dbConn := db.NewDB()
	//deferを使用すると処理を後回しにすることができる.
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	//マイグレートするモデルをアドレスを参照する。
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}