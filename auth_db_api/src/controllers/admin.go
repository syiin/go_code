package controllers

import (
	"encoding/json"
	"fmt"
	"models"
	"net/http"
	"time"
	"utils"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type ErrorResponse struct {
	Err string
}

type error interface {
	Error() string
}

var db = utils.ConnectDB()

func MagaAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("i am here")
}

func TestAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API live and kicking"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	user := &models.Admin{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := FindOne(user.Email, user.Password)
	json.NewEncoder(w).Encode(resp)
}

func FindOne(email, password string) map[string]interface{} {
	user := &models.Admin{}

	if err := db.Table("admins").Where("Email = ?", email).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	tk := &models.Token{
		UserID: user.ID,
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString //Store the token in the response
	resp["user"] = user
	return resp
}

//CreateUser function -- create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {

	user := &models.Admin{}
	json.NewDecoder(r.Body).Decode(user)

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: "Password Encryption  failed",
		}
		json.NewEncoder(w).Encode(err)
	}

	user.Password = string(pass)
	fmt.Println(user)

	createdUser := db.Table("admins").Create(user)
	var errMessage = createdUser.Error

	if createdUser.Error != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(createdUser)
}

//FetchUser function
func FetchUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.Admin
	db.Table("admins").Find(&users)
	// db.Table("admins").Preload("auths").Find(&users)
	json.NewEncoder(w).Encode(users)

	//The below is a superfluous block to verify how contexts work
	userFromCtx := r.Context().Value("user")
	structFromCtx, _ := json.Marshal(userFromCtx)
	fmt.Println("\nFrom FetchUsers(): ")
	fmt.Println(string(structFromCtx))

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.Admin{}
	//get the id from params
	params := mux.Vars(r)
	var id = params["id"]
	//query database for user matching the id and fill the user struct
	db.Table("admins").First(&user, id)
	//decode the body and replace the altered fields in the user struct
	json.NewDecoder(r.Body).Decode(user)
	//save the modified user struct to the table
	db.Table("admins").Save(&user)
	json.NewEncoder(w).Encode(&user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var user models.Admin
	db.Table("admins").First(&user, id)
	db.Table("admins").Delete(&user)
	json.NewEncoder(w).Encode("User deleted")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var user models.Admin
	db.Table("admins").First(&user, id)
	json.NewEncoder(w).Encode(&user)
}
