package controllers

import (
	"book_family/pkg/models"
	"book_family/pkg/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var UserRegistration models.User

func GetUser(w http.ResponseWriter, r *http.Request) {
	userRegistration := models.GetAllUsers()
	res, _ := json.Marshal(userRegistration)
	w.Header().Set("Content-Type", "application/json") // Fix content type here
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		http.Error(w, "Error parsing user ID", http.StatusBadRequest)
		return
	}
	userDetails, _ := models.GetUserById(ID)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "application/json") // Fix content type here
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	RegisterUser := &models.User{}
	utils.ParseBody(r, RegisterUser)
	b := RegisterUser.RegisterUser()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json") // Fix content type here
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		http.Error(w, "Error parsing user ID", http.StatusBadRequest)
		return
	}
	user := models.DeleteUser(ID)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json") // Fix content type here
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var UpdateUser = &models.User{}
	utils.ParseBody(r, UpdateUser)
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		http.Error(w, "Error parsing user ID", http.StatusBadRequest)
		return
	}
	userDetails, db := models.GetUserById(ID)
	if UpdateUser.FirstName != "" {
		userDetails.FirstName = UpdateUser.FirstName
	}
	if UpdateUser.SecondName != "" {
		userDetails.SecondName = UpdateUser.SecondName
	}
	if UpdateUser.Age != 0 {
		userDetails.Age = UpdateUser.Age
	}

	if UpdateUser.City != "" {
		userDetails.City = UpdateUser.City
	}
	if UpdateUser.ZipCode != "" {
		userDetails.ZipCode = UpdateUser.ZipCode
	}
	if UpdateUser.County != "" {
		userDetails.County = UpdateUser.County
	}
	if UpdateUser.State != "" {
		userDetails.State = UpdateUser.State
	}
	if UpdateUser.AcademicQualification != "" {
		userDetails.AcademicQualification = UpdateUser.AcademicQualification
	}
	if UpdateUser.CollegeAttended != "" {
		userDetails.CollegeAttended = UpdateUser.CollegeAttended
	}
	if UpdateUser.StudyArea != "" {
		userDetails.StudyArea = UpdateUser.StudyArea
	}
	db.Save(&userDetails)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "application/json") // Fix content type here
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
