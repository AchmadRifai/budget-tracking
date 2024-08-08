package controllers

import (
	"be/db"
	errorhandlers "be/errorHandlers"
	"be/models"
	myutils "be/myUtils"
	"errors"
	"net/http"

	arrayutils "github.com/AchmadRifai/array-utils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	defer errorhandlers.NormalErrorRest(w, r)
	var users []models.User
	conn := db.DbConnect()
	result := conn.Where("id<>?", 1).Find(&users)
	if result.Error != nil {
		panic(result.Error)
	}
	data := arrayutils.Map(users, func(v models.User, _ int) map[string]interface{} {
		return map[string]interface{}{"id": v.ID, "username": v.UserName, "fullname": v.FullName, "email": v.Email}
	})
	myutils.SendJson(w, map[string]interface{}{"message": "Success", "data": data}, 200)
}

func DelUser(w http.ResponseWriter, r *http.Request) {
	defer errorhandlers.NormalErrorRest(w, r)
	vars := mux.Vars(r)
	conn := db.DbConnect()
	err := conn.Transaction(func(tx *gorm.DB) error {
		var user models.User
		result := tx.Where("id=?", vars["id"]).First(&user)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("user not found")
		}
		result = tx.Delete(&user)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	myutils.SendJson(w, map[string]interface{}{"message": "Success"}, 200)
}
