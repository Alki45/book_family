package models

import (
	"book_family/pkg/config"

	"github.com/jinzhu/gorm"
)

// User defines the structure of the user model
//var db *gorm.DB

type User struct {
	gorm.Model
	FirstName             string `json:"first_name"`
	SecondName            string `json:"second_name"`
	Age                   int    `json:"age"`
	City                  string `json:"city"`
	ZipCode               string `json:"zip_code"`
	County                string `json:"county"`
	State                 string `json:"state"`
	AcademicQualification string `json:"academic_qualification"`
	CollegeAttended       string `json:"college_attended"`
	StudyArea             string `json:"study_area"`
}

func init() {

	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) RegisterUser() *User {
	db.NewRecord(u)
	db.Create(&u)
	return u
}
func GetAllUsers() []User {
	var users []User
	// Fetch all users from the database
	db.Find(&users)
	return users
}

// GetUserById retrieves a user by their ID
func GetUserById(Id int64) (*User, *gorm.DB) {
	var user User
	// Find the user by ID
	result := db.Where("ID = ?", Id).Find(&user)
	return &user, result
}

// DeleteUser deletes a user by their ID
func DeleteUser(ID int64) User {
	var user User
	// Delete the user with the given ID
	db.Where("ID = ?", ID).Delete(&user)
	return user
}
