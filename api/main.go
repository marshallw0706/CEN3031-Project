package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// struct for user accounts and files
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Files    []File `json:"files" gorm:"foreignkey:OwnerID"`
}

type File struct {
	gorm.Model
	Filename  string `json:"filename" gorm:"not null"`
	Size      int64  `json:"size" gorm:"not null"`
	Type      string `json:"type" gorm:"not null"`
	OwnerID   string `json:"owner_id" gorm:"not null"`
	CreatedAt int64  `json:"created_at" gorm:"not null"`
}

// Define a global variable for the database connection
var db *gorm.DB

// Initialize the database connection
func initDB() {

	var err error
	dns := "root:@tcp(127.0.0.1:3306)/muse?charset=utf8&parseTime=true"
	// Connect to the database
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Create a new database if one does not exist
	err = db.Exec("CREATE DATABASE IF NOT EXISTS muse").Error
	if err != nil {
		panic("failed to create database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})
	db.AutoMigrate(&File{})
}

// Initialize the routes and handlers
func initRouter() {
	r := mux.NewRouter()

	// handlers for user routes
	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", getUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", deleteUser).Methods("DELETE")

	// handlers for file routes
	r.HandleFunc("/api/users/{id}/files", getFiles).Methods("GET")
	r.HandleFunc("/api/users/{id}/files", createFile).Methods("POST")
	r.HandleFunc("/api/users/{id}/files/{fid}", updateFile).Methods("PUT")
	r.HandleFunc("/api/users/{id}/files/{fid}", deleteFile).Methods("DELETE")

	// Start listening on port 5000 using http package
	fmt.Println("Starting server on port 5000...")
	log.Fatal(http.ListenAndServe(":5000", r))
}

// handler to create a new user account
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User

	// Decode the request body into the user struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Validate the user input (e.g. check if username or password is empty)
	if user.Username == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "username or password cannot be empty"})
		return
	}

	// Create the user in the database
	err = db.Create(&user).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "error creating user"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// handler to get all user accounts
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []User

	err := db.Find(&users).Error // Find all users in the database

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "error getting users"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users) // Encode and send back all users as JSON response
}

// handler for updating a user account
func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User

	// Decode the request body into the user struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the user ID from the URL parameter (e.g. /users/{id})
	params := mux.Vars(r)
	userID := params["id"]

	// Update the user in the database
	err = db.Model(&user).Where("id = ?", userID).Update("username", user.Username).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "error updating user"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user) // Encode and send back the user as JSON response
}

// handler for deleting a user account
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the user ID from the URL parameter (e.g. /users/{id})
	params := mux.Vars(r)
	userID := params["id"]

	// Delete the user from the database
	err := db.Where("id = ?", userID).Delete(&User{}).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "error deleting user"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "user deleted"})
}

// handler to create a new file for a given user account
func createFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var file File

	// Decode the request body into the file struct
	err := json.NewDecoder(r.Body).Decode(&file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Validate the file input (e.g. check if filename or size is empty)
	if file.Filename == "" || file.Size == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "filename or size cannot be empty"})
		return
	}

	// Get the owner ID from the URL parameter (e.g. /users/{id}/files)
	params := mux.Vars(r)
	ownerID := params["id"]

	// Set the owner ID and created at fields for the file struct
	file.OwnerID = ownerID

	file.CreatedAt = time.Now().Unix()

	// Create the file in the database
	err = db.Create(&file).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "error creating file"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(file) // Encode and send back the file as JSON response
}

// handler to get all files for a given user account
func getFiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var files []File

	// Get the owner ID from the URL parameter (e.g. /users/{id}/files)
	params := mux.Vars(r)
	ownerID := params["id"]

	err := db.Where("owner_id = ?", ownerID).Find(&files).Error // Find all files that belong to the owner in the database

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "error getting files"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(files) // Encode and send back all files as JSON response
}

// handler to update a file for a given user account
func updateFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var file File

	// Decode the request body into the file struct
	err := json.NewDecoder(r.Body).Decode(&file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the owner ID and file ID from the URL parameter (e.g. /users/{id}/files/{fid})
	params := mux.Vars(r)
	ownerID := params["id"]
	fileID := params["fid"]

	err = db.Where("owner_id = ? AND id = ?", ownerID, fileID).First(&file).Error // Find the file that belongs to the owner and has the given ID in the database

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "file not found"})
		return
	}

	// Update only non-empty fields of the file struct in the database
	db.Model(&file).Updates(file)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(file) // Encode and send back updated file as JSON response
}

// handler to delete a file for a given user account
func deleteFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var file File

	// Get the owner ID and file ID from URL parameter (e.g. /users/{id}/files/{fid})
	params := mux.Vars(r)
	ownerID := params["id"]
	fileID := params["fid"]

	err := db.Where("owner_id = ? AND id = ?", ownerID, fileID).First(&file).Error // Find the file that belongs to owner and has given ID in database

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "file not found"})
		return
	}

	// Delete only non-empty fields of the file struct in the database
	db.Delete(&file)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "file deleted"}) // Encode and send back a confirmation message as JSON response
}

func main() {
	initDB()     // Initialize database connection
	initRouter() // Initialize router and handlers
}
