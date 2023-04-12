package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initializeRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/api/users", GetUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/api/users", CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}", DeleteUser).Methods("DELETE")

	router.HandleFunc("/api/users/{id}/profile", UpdateProfileInfo).Methods("PUT")
	router.HandleFunc("/api/users/{id}/profile", GetProfileInfo).Methods("GET")

	router.HandleFunc("/api/users/{id}/follow/{fid}", addFollower).Methods("POST")
	router.HandleFunc("/api/users/{id}/unfollow/{fid}", removeFollower).Methods("DELETE")
	router.HandleFunc("/api/users/{id}/following", getFollowingUsers).Methods("GET")
	router.HandleFunc("/api/users/{uid}/like/{id}/{fid}", likeFile).Methods("POST")
	router.HandleFunc("/api/users/{uid}/unlike/{id}/{fid}", unlikeFile).Methods("DELETE")
	router.HandleFunc("/api/users/{id}/files/{fid}/likedby", getLikedByUsers).Methods("GET")

	router.HandleFunc("/api/users/{id}/files", getFiles).Methods("GET")
	router.HandleFunc("/api/users/{id}/files", createFile).Methods("POST")
	router.HandleFunc("/api/users/{id}/files/upload", uploadFile).Methods("POST")
	router.HandleFunc("/api/users/{id}/files/{fid}", updateFile).Methods("PUT")
	router.HandleFunc("/api/users/{id}/files/{fid}", deleteFile).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))
}

func main() {
	InitialMigration()
	initializeRouter()
}

var DB *gorm.DB
var err error

const DSN = "root:@tcp(localhost:3306)/sys?charset=utf8&parseTime=true"

type User struct {
	gorm.Model
	Username    string        `json:"username"`
	Password    string        `json:"password"`
	Files       []File        `json:"files" gorm:"foreignkey:OwnerID"`
	ProfileInfo ProfileStruct `json:"profileinfo"  gorm:"foreignkey:OwnerID"`
	Following   []User        `json:"following" gorm:"many2many:user_followers"`
}

type File struct {
	gorm.Model
	Filename  string `json:"filename" gorm:"not null"`
	Size      int64  `json:"size" gorm:"not null"`
	Type      string `json:"type" gorm:"not null"`
	OwnerID   string `json:"owner_id" gorm:"not null"`
	CreatedAt int64  `json:"created_at" gorm:"not null"`
	Data      []byte `json:"data" gorm:"not null"`
	Likes     int64  `json:"likes" gorm:"not null"`
	LikedBy   []User `json:"likedby" gorm:"many2many:file_likes"`
}

type ProfileStruct struct {
	OwnerID     uint   `json:"owner_id" gorm:"not null"`
	Name        string `json:"name"`
	JobTitle    string `json:"jobtitle"`
	Description string `json:"description"`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&User{}, &File{}, &ProfileStruct{})
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)
	DB.Create(&user)
	profile := ProfileStruct{
		OwnerID: user.ID,
	}
	DB.Create(&profile)

	user.ProfileInfo = profile

	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(w).Encode(user)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("The User has been deleted")
}

func UpdateProfileInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var profile ProfileStruct
	DB.Where("owner_id = ?", params["id"]).First(&profile)

	var updatedProfile ProfileStruct
	if err := json.NewDecoder(r.Body).Decode(&updatedProfile); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if updatedProfile.Name != "" {
		profile.Name = updatedProfile.Name
	}
	if updatedProfile.JobTitle != "" {
		profile.JobTitle = updatedProfile.JobTitle
	}
	if updatedProfile.Description != "" {
		profile.Description = updatedProfile.Description
	}

	DB.Model(&profile).Where("owner_id = ?", profile.OwnerID).Updates(profile)
	json.NewEncoder(w).Encode(profile)
}

func GetProfileInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var profile ProfileStruct
	result := DB.Where("owner_id = ?", params["id"]).First(&profile)

	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "profile not found"})
		return
	}

	json.NewEncoder(w).Encode(profile)
}

func addFollower(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var user, follower User
	DB.First(&user, params["id"])
	DB.First(&follower, params["fid"])

	if user.ID == 0 || follower.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "user or follower not found"})
		return
	}

	err := DB.Model(&user).Association("Following").Append(&follower)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "error adding follower"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func removeFollower(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var user, follower User
	DB.First(&user, params["id"])
	DB.First(&follower, params["fid"])

	if user.ID == 0 || follower.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "user or follower not found"})
		return
	}

	err := DB.Model(&user).Association("Following").Delete(&follower)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "error removing follower"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func likeFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var user User
	var file File
	DB.First(&user, params["uid"])
	ownerID := params["id"]
	fileID := params["fid"]

	result := DB.Where("owner_id = ? AND id = ?", ownerID, fileID).First(&file)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "file not found"})
		return
	}

	err := DB.Model(&file).Association("LikedBy").Append(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "error liking file"})
		return
	}

	file.Likes++
	DB.Save(&file)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(file)
}

func unlikeFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var user User
	var file File
	DB.First(&user, params["uid"])
	ownerID := params["id"]
	fileID := params["fid"]
	result := DB.Where("owner_id = ? AND id = ?", ownerID, fileID).First(&file)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "file not found"})
		return
	}

	err := DB.Model(&file).Association("LikedBy").Delete(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "error unliking file"})
		return
	}

	file.Likes--
	DB.Save(&file)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(file)
}

func getFollowingUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	DB.Preload("Following").First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "user not found"})
		return
	}

	var followingUsers []User
	for _, followingUserID := range user.Following {
		var followingUser User
		DB.First(&followingUser, followingUserID)
		followingUsers = append(followingUsers, followingUser)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(followingUsers)
}

func getLikedByUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var file File
	ownerID := params["id"]
	fileID := params["fid"]
	DB.Preload("LikedBy").Where("owner_id = ? AND id = ?", ownerID, fileID).First(&file)
	result := DB.Where("owner_id = ? AND id = ?", ownerID, fileID).First(&file)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "file not found"})
		return
	}

	var likedByUsers []User
	for _, likedByUserID := range file.LikedBy {
		var likedByUser User
		DB.First(&likedByUser, likedByUserID)
		if likedByUser.ID != 0 {
			likedByUsers = append(likedByUsers, likedByUser)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(likedByUsers)
}

// handler for uploading a file to an account
func uploadFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the user ID from the URL parameter (e.g. /users/{id}/files)
	params := mux.Vars(r)
	userID := params["id"]

	// Get the file from the request body
	file, header, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "error getting file"})
		return
	}
	defer file.Close()

	// Create a new file with the uploaded files properties
	newFile := File{
		Filename:  header.Filename,
		Size:      header.Size,
		Type:      header.Header.Get("Content-Type"),
		OwnerID:   userID,
		CreatedAt: time.Now().Unix(),
	}

	newFile.Data = make([]byte, newFile.Size) // Create a byte slice with the size of the file
	_, err = file.Read(newFile.Data)          // Read the file into the byte slice
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "error reading file data"})
		return
	}

	// Create the file in the database
	err = DB.Create(&newFile).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "error creating file: " + err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newFile)
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
	err = DB.Create(&file).Error
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

	err := DB.Where("owner_id = ?", ownerID).Find(&files).Error // Find all files that belong to the owner in the database

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

	// Get the owner ID and file ID from the URL parameter (e.g. /users/{id}/files/{fid})
	params := mux.Vars(r)
	ownerID := params["id"]
	fileID := params["fid"]

	var file File
	result := DB.Where("owner_id = ? AND id = ?", ownerID, fileID).First(&file) // Find the file that belongs to the owner and has the given ID in the database

	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "file not found"})
		return
	}

	// Decode the request body into the file struct and update only non-empty fields in the database
	err := json.NewDecoder(r.Body).Decode(&file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	DB.Save(&file)

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

	err := DB.Where("owner_id = ? AND id = ?", ownerID, fileID).First(&file).Error // Find the file that belongs to owner and has given ID in database

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "file not found"})
		return
	}

	// Delete only non-empty fields of the file struct in the database
	DB.Delete(&file)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "file deleted"}) // Encode and send back a confirmation message as JSON response
}
