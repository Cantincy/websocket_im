package entity

type User struct {
	UserId string `json:"user_id" gorm:"column:userId;primary_key"`
	Pwd    string `json:"pwd" gorm:"column:pwd;not null"`
}

func (u User) TableName() string {
	return "user_tb"
}
