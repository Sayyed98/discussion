package auth

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Id        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	Phone     int    `json:"phone" db:"phone"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
	DeletedAt string `json:"deleted_at" db:"deleted_at"`
}

func Signup(w http.ResponseWriter, r *http.Request) {

	// parsing request
	req := User{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Fatal("error in request parsing")
	}

}
