package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
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

func TestGetUser(t *testing.T) {
	InitialMigration()
	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}", GetUser).Methods("GET")

	user1 := &User{Username: "testuser1", Password: "testpass1"}
	DB.Create(&user1)
	var idUser User
	DB.Where("username = ?", "testuser1").First(&idUser)
	id := idUser.ID

	req, err := http.NewRequest("GET", "/api/users/"+strconv.Itoa(int(id)), nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}

	var user User
	json.Unmarshal(rr.Body.Bytes(), &user)
	if user.Username != "testuser1" {
		t.Errorf("handler returned wrong username: got %v want %v", user.Username, "testuser1")
	}

	DB.Where("username = ?", "testuser1").Delete(&user1)
}

func TestUploadFile(t *testing.T) {
	// Initialize the database for testing
	InitialMigration()

	// Create a test user
	user := User{
		Username: "testuser",
		Password: "testpassword",
	}
	DB.Create(&user)

	// Prepare the request body
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	file, err := os.Open("./testfile.txt") // Make sure you have a testfile.txt in the same directory
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	part, err := writer.CreateFormFile("file", "testfile.txt")
	if err != nil {
		t.Fatal(err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatal(err)
	}
	err = writer.Close()
	if err != nil {
		t.Fatal(err)
	}

	// Create a test request
	req, err := http.NewRequest("POST", "/api/users/"+strconv.Itoa(int(user.ID))+"/files/upload", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Execute the request using httptest
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/files/upload", uploadFile).Methods("POST")
	router.ServeHTTP(rr, req)

	// Check the response status code and content
	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, rr.Code)
	}

	var uploadedFile File
	err = json.Unmarshal(rr.Body.Bytes(), &uploadedFile)
	if err != nil {
		t.Fatal(err)
	}

	if uploadedFile.Filename != "testfile.txt" {
		t.Errorf("Expected filename %s, got %s", "testfile.txt", uploadedFile.Filename)
	}

	if uploadedFile.OwnerID != strconv.FormatUint(uint64(user.ID), 10) {
		t.Errorf("Expected owner ID %s, got %s", strconv.FormatUint(uint64(user.ID), 10), uploadedFile.OwnerID)
	}

	// Clean up test data
	DB.Where("id = ?", user.ID).Delete(&user)
	DB.Where("id = ?", uploadedFile.ID).Delete(&uploadedFile)
}

func TestCreateFile(t *testing.T) {
	InitialMigration()
	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/files", createFile).Methods("POST")
	user := &User{Username: "testuser6", Password: "testpass6"}
	DB.Create(&user)
	var idUser User
	DB.Where("username = ?", "testuser6").First(&idUser)
	id := idUser.ID

	newFile := &File{
		Filename:  "testfile",
		Size:      1024,
		Type:      "text/plain",
		OwnerID:   strconv.Itoa(int(id)),
		CreatedAt: time.Now().Unix(),
		Data:      []byte("test file data"),
	}

	jsonFile, err := json.Marshal(newFile)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/api/users/"+strconv.Itoa(int(id))+"/files", bytes.NewBuffer(jsonFile))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if rr.Code != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusCreated)
	}

	var createdFile File
	json.Unmarshal(rr.Body.Bytes(), &createdFile)

	if createdFile.Filename != "testfile" {
		t.Errorf("handler returned wrong filename: got %v want %v",
			createdFile.Filename, "testfile")
	}

	DB.Where("filename = ?", "testfile").Delete(&createdFile)
	DB.Where("username = ?", "testuser6").Delete(&user)
}

func TestGetFiles(t *testing.T) {
	InitialMigration()

	// Create a new user to own the files
	user := User{
		Username: "testuser",
		Password: "testpassword",
	}
	DB.Create(&user)
	fmt.Printf("TestGetFiles - User ID: %d\n", user.ID)

	// Create two new files to belong to the user
	file1 := File{
		Filename:  "testfile1.txt",
		Size:      1024,
		Type:      "text/plain",
		OwnerID:   strconv.Itoa(int(user.ID)),
		CreatedAt: time.Now().Unix(),
		Data:      []byte(strings.Repeat("a", 1024)),
	}
	DB.Create(&file1)

	file2 := File{
		Filename:  "testfile2.txt",
		Size:      2048,
		Type:      "text/plain",
		OwnerID:   strconv.Itoa(int(user.ID)),
		CreatedAt: time.Now().Unix(),
		Data:      []byte(strings.Repeat("b", 2048)),
	}
	DB.Create(&file2)

	// Create a new HTTP request to get the user's files
	req, err := http.NewRequest("GET", "/api/users/"+strconv.Itoa(int(user.ID))+"/files", nil)
	if err != nil {
		t.Fatal(err)
	}

	// execute the request using httptest
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/files", getFiles).Methods("GET")
	router.ServeHTTP(rr, req)

	// Check the status code returned by the server
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Unmarshal the response body into a slice of files
	var files []File
	err = json.Unmarshal(rr.Body.Bytes(), &files)
	if err != nil {
		t.Errorf("could not unmarshal response body: %v", err)
	}

	// Check that the correct number of files was returned
	if len(files) != 2 {
		t.Errorf("handler returned wrong number of files: got %v want %v", len(files), 2)
	}

	// Check that the correct files were returned
	fileIDSet := map[uint]bool{
		files[0].ID: true,
		files[1].ID: true,
	}

	if !fileIDSet[file1.ID] || !fileIDSet[file2.ID] {
		t.Errorf("handler returned wrong files")
	}

	// Clean up by deleting the test user and files from the database
	DB.Delete(&file1)
	DB.Delete(&file2)
	DB.Delete(&user)
}

func TestUpdateFile(t *testing.T) {
	InitialMigration()
	// Create a new user to own the file
	user := User{
		Username: "testuser",
		Password: "testpassword",
	}
	DB.Create(&user)

	// Create a file owned by the user
	file := File{
		Filename:  "testfile.txt",
		Size:      1024,
		Type:      "text/plain",
		OwnerID:   strconv.Itoa(int(user.ID)),
		CreatedAt: time.Now().Unix(),
		Data:      []byte(strings.Repeat("a", 1024)),
	}
	DB.Create(&file)

	// Update the file's filename and type
	file.Filename = "updatedfile.txt"
	file.Type = "text/html"

	// Encode the file struct as JSON and create a new HTTP request to update the file
	payload, err := json.Marshal(file)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("PUT", "/api/users/"+strconv.Itoa(int(user.ID))+"/files/"+strconv.Itoa(int(file.ID)), bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/files/{fid}", updateFile).Methods("PUT")
	router.ServeHTTP(rr, req)

	// Check the status code returned by the server
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Unmarshal the response body into a file struct
	var updatedFile File
	err = json.Unmarshal(rr.Body.Bytes(), &updatedFile)
	if err != nil {
		t.Errorf("could not unmarshal response body: %v", err)
	}

	// Check that the filename and type of the file were updated
	if updatedFile.Filename != "updatedfile.txt" || updatedFile.Type != "text/html" {
		t.Errorf("file was not updated: got %+v want %+v", updatedFile, file)
	}

	// Clean up by deleting the test user and file from the database
	DB.Delete(&file)
	DB.Delete(&user)
}

func TestDeleteFile(t *testing.T) {
	InitialMigration()
	// Create a new user to own the file
	user := User{
		Username: "testuser",
		Password: "testpassword",
	}
	DB.Create(&user)

	// Create a file owned by the user
	file := File{
		Filename:  "testfile.txt",
		Size:      1024,
		Type:      "text/plain",
		OwnerID:   strconv.Itoa(int(user.ID)),
		CreatedAt: time.Now().Unix(),
		Data:      []byte(strings.Repeat("a", 1024)),
	}
	DB.Create(&file)

	// Create a new HTTP request to delete the file
	req, err := http.NewRequest("DELETE", "/api/users/"+strconv.Itoa(int(user.ID))+"/files/"+strconv.Itoa(int(file.ID)), nil)
	if err != nil {
		t.Fatal(err)
	}

	// execute the request using httptest
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/files/{fid}", deleteFile).Methods("DELETE")
	router.ServeHTTP(rr, req)

	// Check the status code returned by the server
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check that the file was deleted from the database
	var deletedFile File
	err = DB.Where("id = ?", file.ID).First(&deletedFile).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("file was not deleted from the database")
	}

	// Clean up by deleting the test user from the database
	DB.Delete(&user)
}

func TestGetComments(t *testing.T) {
	InitialMigration()
	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}/files/{fid}/comments", getComments).Methods("GET")

	// Create test data
	user := &User{Username: "testuser", Password: "testpass"}
	DB.Create(user)

	file := &File{
		Filename:    "testfile",
		Size:        1024,
		Type:        "text/plain",
		OwnerID:     strconv.Itoa(int(user.ID)),
		CreatedAt:   time.Now().Unix(),
		Data:        []byte("Test data"),
		Likes:       0,
		Description: "Test description",
	}
	DB.Create(file)

	comment := &Comment{
		Content: "Test comment",
		UserID:  user.ID,
		FileID:  file.ID,
	}
	DB.Create(comment)

	// Test the endpoint
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/users/%d/files/%d/comments", user.ID, file.ID), nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	var comments []Comment
	err = json.Unmarshal(rr.Body.Bytes(), &comments)
	if err != nil {
		t.Fatal(err)
	}

	if len(comments) != 1 || comments[0].Content != comment.Content {
		t.Errorf("handler returned unexpected comments: got %v want %v",
			comments, []Comment{*comment})
	}

	// Clean up test data
	DB.Delete(&comment)
	DB.Delete(&file)
	DB.Delete(&user)
}
