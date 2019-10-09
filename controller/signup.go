package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"helpin/model"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/lib/pq"
)

//SignUp Controllers
func (cntrl Controllers) SignUp(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		var err error
		json.NewDecoder(c.Request.Body).Decode(&user)
		// x, _ := ioutil.ReadAll(user)
		user.Token, err = GenerateToken(user.Username)
		logFatal(err)
		if err != nil {
			SetErrorResponse(c, http.StatusInternalServerError, "failed to generate token", http.StatusInternalServerError)
			return
		}
		user.Password, err = HashPassword(user.Password)
		logFatal(err)
		if err != nil {
			SetErrorResponse(c, http.StatusInternalServerError, "failed to hashing password", http.StatusInternalServerError)
			return
		}

		err = db.QueryRow("insert into user_login(username,password,token) values "+
			"($1,$2,$3) returning user_id", user.Username, user.Password, user.Token).Scan(&user.ID)

		if err != nil {
			fmt.Println(err)
			pqErr := err.(*pq.Error)
			if pqErr.Code == "23505" {
				SetErrorResponse(c, http.StatusOK, "Username already exist.", 400)
				return
			}

			SetErrorResponse(c, http.StatusInternalServerError, pqErr.Message, http.StatusInternalServerError)
			return
		}
		SetResponse(c, http.StatusOK, user)
		return
	}
}
