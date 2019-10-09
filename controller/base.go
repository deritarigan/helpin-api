package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"helpin/model"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/heroku/appsention-api/models"
	"golang.org/x/crypto/bcrypt"
)

//Controllers to initiate controllers
type Controllers struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func setControllerDB(database *sql.DB) {
	// db = database
}

//SetErrorResponse
func SetResponse(c *gin.Context, code int, responseData interface{}) {
	var metaModel model.MetaModel
	metaModel.Message = "Success"
	metaModel.Code = code
	metaModel.Status = false
	c.JSON(code, gin.H{
		"data": responseData,
		"meta": metaModel,
	})
}

//SetErrorResponse
func SetErrorResponse(c *gin.Context, code int, message string, metaCode int) {
	var metaModel model.MetaModel
	var empty struct{}
	metaModel.Message = message
	metaModel.Code = metaCode
	metaModel.Status = false
	c.JSON(code, gin.H{
		"data": empty,
		"meta": metaModel,
	})
}

func responseJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func getEmptyData() *model.Object {
	var object = new(model.Object)
	return object
}

func checkPassword(w http.ResponseWriter, db *sql.DB, passwordModel models.ChangePassword) bool {
	var currPassword string
	var errorModel models.ErrorModel
	err := db.QueryRow("select password from users where user_id =$1",
		passwordModel.ID).Scan(&currPassword)
	logFatal(err)

	err = bcrypt.CompareHashAndPassword([]byte(currPassword), []byte(passwordModel.Password))
	if err != nil {
		errorModel.Message = "Invalid password"
		// respondWithError(w, http.StatusBadRequest, errorModel)
		return false
	}

	return true
}

//HashPassword is a function to hasing the given password
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	return string(hash), err
}

//GenerateToken is function to generate JWT
func GenerateToken(user string) (string, error) {
	var err error
	secret := "secret"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user,
		"iss":      strconv.FormatInt(GetCurrentTimeMillis(), 10),
	})
	tokenString, err := token.SignedString([]byte(secret))

	return tokenString, err
}

//GetCurrentTimeMillis to generate time in epoch
func GetCurrentTimeMillis() int64 {
	nanos := time.Now().UnixNano()

	return nanos / 1000000
}

//CompareInsensitive is a function to comparing string value
func CompareInsensitive(a, b string) bool {
	// a quick optimization. If the two strings have a different
	// length then they certainly are not the same
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		// if the characters already match then we don't need to
		// alter their case. We can continue to the next rune
		if a[i] == b[i] {
			continue
		} else {
			return false
		}
	}
	// The string length has been traversed without a mismatch
	// therefore the two match
	return true
}

func verifyToken(r *http.Request) (string, jwt.MapClaims) {
	authHeader := r.Header.Get("Authorization")
	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) == 2 {
		authToken := bearerToken[1]

		token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte("secret"), nil
		})
		if err != nil {
			return "", nil
		}
		logFatal(err)
		if token.Valid {
			jwtMap := token.Claims.(jwt.MapClaims)
			return authToken, jwtMap
		}
	}
	return "", nil
}
