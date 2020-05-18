package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gofrs/uuid"
	postgresql "github.com/xragon/simplelogin/internal/postgres"
	"golang.org/x/crypto/bcrypt"
)

// type User struct {
// 	username string
// 	password string
// }

func CreateUser(w http.ResponseWriter, req *http.Request) {
	var u postgresql.User

	err := json.NewDecoder(req.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	u.Password = string(hash)

	u.ID, _ = uuid.NewV4()

	DB, _ := postgresql.NewStore()

	DB.WriteRecord(u)

}

func Login(w http.ResponseWriter, req *http.Request) {
	var u postgresql.User

	err := json.NewDecoder(req.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	DB, _ := postgresql.NewStore()

	dbuser, err := DB.GetUser(u.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbuser.Password), []byte(u.Password))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	fmt.Fprint(w, "Success")
}

func getData(w http.ResponseWriter, req *http.Request) {

}

func main() {

	http.HandleFunc("/create", CreateUser)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/getdata", getData)

	http.ListenAndServe(":8090", nil)
}
