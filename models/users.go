package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Users struct {
	db *gorm.DB
}

type User struct {
	gorm.Model
	Name string
}

func NewUsers() *Users {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/node_rest_api?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	db.AutoMigrate(&User{})

	return &Users{db}
}

func (m *Users) Create(name string) (*User, error) {
	user := &User{Name: name}
	if err := m.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (m *Users) Delete(id int) (int64, error) {
	result := m.db.Unscoped().Delete(&User{}, "id = ?", id)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

func (m *Users) Selete(id int) (User, error) {
	user := User{}
	err := m.db.Find(&user, "id = ?", id).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (m *Users) SeleteAll() ([]User, error) {
	users := []User{}
	m.db.Find(&users)

	return users, nil
}
