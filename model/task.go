package model

import "time"

//taskについての構造体を定義する。
type Task struct{
	//taskのidの定義。主キーである。
	ID uint `json:"id" gorm:"primaryKey"`

	//titleはnot null
	Title string `json:"title" gorm:"not null"`

	//作成時と更新時の時間を記録する.
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	//user.goで作成したUser構造体を格納する.
	//constrain:OnDelete:CASCADEにすることでUser情報が削除された際に紐づいているtaskを全て削除する。
	User User `json:"user" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`

	//UserIdはnot null
	UserId uint `json:"user_id" gorm:"not null"`
}

//クライアントからGETメソッドで求められた際のTaskResponseを定義する.
type TaskResponse struct{
	ID uint `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

