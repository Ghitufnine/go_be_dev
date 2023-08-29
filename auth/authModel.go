package auth

import "go_be_dev/database"

type User struct {
	Id       int    `gorm:"column:id" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Name     string `gorm:"column:name" json:"name"`
}

func PostDoLoginQuery(username string, password string) []User {
	result := []User{}

	database.DBConn.Unscoped().
		Table("user").
		Where("username = ?", username).
		Where("password = ?", password).
		Find(&result)

	return result
}
