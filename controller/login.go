package controller

import (
	"database/sql"
	"encoding/json"
	"helpin/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//Login Controllers
func (cntrl Controllers) Login(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		var err error
		json.NewDecoder(c.Request.Body).Decode(&user)
		userPassword := user.Password
		row := db.QueryRow("select * from user_login where username =$1", user.Username)
		err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Token)
		if err != nil {
			if err == sql.ErrNoRows {
				SetErrorResponse(c, http.StatusOK, "User doesnt exist", 404)
				return
			}
			logFatal(err)
		}
		hashedPassword := user.Password
		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userPassword))

		if err != nil {
			SetErrorResponse(c, http.StatusOK, "Invalid password.", 400)
			return
		}

		token, err := GenerateToken(user.Username)
		logFatal(err)

		err = db.QueryRow("update user_login set token =$1 where user_id =$2 returning token,user_id;", &token,
			user.ID).Scan(&user.Token, &user.ID)
		logFatal(err)

		SetResponse(c, http.StatusOK, user)
	}
}
