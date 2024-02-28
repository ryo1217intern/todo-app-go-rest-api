// パッケージはディレクトリ毎に分類する
package model

import "time"

//Userの構造体の定義
type User struct {
	//uintは負の値を取らない整数型, jsonの値は基本的に小文字にする, IDは主キーのためprimaryKeyい設定する.
	ID uint `json:"id" gorm:"primaryKey"`

	//任意のEmailはデータベース上に一つで一つでなくてはならないのでuniqueとする.
	Email string `json:"email" gorm:"unique"`

	Password string `json:"password"`

	//作成時とアップデートした時間をtimeライブラリを使用して定義する.
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time .Time `json:"updated_at"`
}

//クライアントに渡すデータの構造体の定義.
type UserResponse struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Email string `json:"email" gorm:"unique"`
}