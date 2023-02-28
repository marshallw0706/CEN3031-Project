package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
)

func TestUpdateUser(t *testing.T) {
	InitialMigration()
	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}", UpdateUser).Methods("PUT")

	user1 := &User{Username: "testuser3", Password: "testpass3"}
	DB.Create(&user1)
	var idUser User
	DB.Where("username = ?", "testuser3").First(&idUser)
	id := idUser.ID

	var jsonStr string = "{\"username\": \"testuser3new\"}"
	req, err := http.NewRequest("PUT", "/api/users/"+strconv.Itoa(int(id)), bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	var updatedUser User
	DB.Find(&updatedUser, id)

	if updatedUser.Username != "testuser3new" {
		t.Errorf("handler returned wrong username: got %v want %v",
			updatedUser.Username, "testuser3new")
	}

	DB.Where("username = ?", "testuser3new").Delete(&user1)
}

func TestCreateUser(t *testing.T) {
	InitialMigration()
	router := mux.NewRouter()
	router.HandleFunc("/api/users", CreateUser).Methods("POST")

	user := &User{Username: "testuser", Password: "testpass"}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	var dbUser User
	if err := DB.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
		t.Fatal(err)
	}
	DB.Where("username = ?", "testuser").Delete(&user)

}

func TestGetUsers(t *testing.T) {
	InitialMigration()
	router := mux.NewRouter()
	router.HandleFunc("/api/users", GetUsers).Methods("GET")

	user1 := &User{Username: "testuser1", Password: "testpass1"}
	user2 := &User{Username: "testuser2", Password: "testpass2"}
	DB.Create(&user1)
	DB.Create(&user2)

	req, err := http.NewRequest("GET", "/api/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	var users []User
	json.Unmarshal(rr.Body.Bytes(), &users)
	if len(users) != 2 {
		t.Errorf("handler returned wrong number of users: got %v want %v",
			len(users), 2)
	}

	DB.Where("username = ?", "testuser1").Delete(&user1)
	DB.Where("username = ?", "testuser2").Delete(&user2)
}
