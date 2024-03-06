// パッケージ名は所属しているディレクトリ名
package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//データベースを作成する関数を作成する
//パッケージの中で定義されているDBの構造体のアドレスを関数の返り値としてgormのDBのポイント型を定義する.
func NewDB() *gorm.DB {
	//環境変数を読み込むための処理.
	//第一条件に.envファイルのGO_ENVの変数の値がdevであった際に実行する.
  if os.Getenv("GO_ENV") == "dev" {
	  //関数では基本的にreturnの値はエラーになっている。処理はあくまで副産物的扱い。
		//errの値がnilではない。つまりエラーがあった際はlogにエラーを出力する.
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}
	//データベースのURLをsprintfを使用して代入する.
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	//dbとerr変数にgorm.Openを行う
	//gorm.Openの引数はpostgresのURLとgorm.Configを渡す.
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	//err処理を記入する。
	if err != nil {
		log.Fatalln(err)
	}
	//正常にプログラムが動作した際にlogに出力する.
	fmt.Println("Connected")
	//returnする値は作成したdb
	return db
}

//データベースをクローズする関数を作成する.
//引数はdbでありgorm.DBのポイント型をとる.
func CloseDB(db *gorm.DB) {
	//.DB()メソッドは返り値に*sql.DB, errorを返すが今回エラーは使わないので＿で破棄する。
	sqlDB, _ := db.DB()
	//Close処理
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
