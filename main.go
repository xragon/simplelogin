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

	// hash := md5.Sum([]byte(u.password))
	// u.password = hex.EncodeToString(hash[:])

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	u.Password = string(hash)

	u.ID, _ = uuid.NewV4()

	DB, _ := postgresql.NewStore()

	DB.WriteRecord(u)

}

func login(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	u := q.Get("username")
	p := q.Get("password")

	fmt.Fprintf(w, "hello %s, %s\n", u, p)
}

func getData(w http.ResponseWriter, req *http.Request) {

}

func main() {

	http.HandleFunc("/create", CreateUser)
	http.HandleFunc("/login", login)
	http.HandleFunc("/getdata", getData)

	http.ListenAndServe(":8090", nil)
}
