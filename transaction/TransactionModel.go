package transaction

import "go_be_dev/database"

type User struct {
	Id       int    `gorm:"column:id" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Name     string `gorm:"column:name" json:"name"`
}

func GetUserInfoByUsername(username string) []User {
	result := []User{}

	database.DBConn.Unscoped().
		Table("user").
		Where("username = ?", username).
		Find(&result)

	return result
}

type DataTransaction struct {
	UserId    int     `gorm:"column:user_id"`
	IpAddress string  `gorm:"column:ip_address"`
	Latitude  float64 `gorm:"column:latitude"`
	Longitude float64 `gorm:"column:longitude"`
	ClockIn   string  `gorm:"column:clock_in"`
}
type DataTransaction2 struct {
	UserId    int     `gorm:"column:user_id"`
	IpAddress string  `gorm:"column:ip_address"`
	Latitude  float64 `gorm:"column:latitude"`
	Longitude float64 `gorm:"column:longitude"`
	ClockOut  string  `gorm:"column:clock_out"`
}

func TransactionClockIn(data DataTransaction) error {
	tx := database.DBConn.Begin()
	if err := tx.
		Table("tr_clock_in").
		Create(&data).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func TransactionClockOut(data DataTransaction2) error {
	tx := database.DBConn.Begin()
	if err := tx.
		Table("tr_clock_out").
		Create(&data).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
