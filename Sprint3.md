# Front-End:

In this sprint, we worked on getting the api and client to work together and focused on making sure that the 

login/sign up page worked together and that a user could create an account. In this sprint we also edited the profile

page so that the user could edit it and put in their own personalized informaiton. The user will also have the capability to 

upload there own profile image. We also edited the explore page to show how it will look like once the user logs in.

This page was updated to hold posts from different artists that the user does not follow. Each post shown on the explore

page is next to one another and is a button so that the user can click on it. Once the user clicks on a photo they will be sent to the profile of the artist. 

There is also a caption box for the artist to write a small description. The explore page will also have a search bar so the user can look up a certain 

artist or a genre of song. The logo on the home page was also updated. The post page was updated so

that the user can choose a file and upload it. Once uploaded the user will either be able to see

the image or be able to play the song due to the audio bar that created. We also added different fonts to our project so

that we can use them in the future. 

Unit Tests - Below is the unit tests that were created for the front end

  async checkUser()
  {
    const users$ = await this.httpClient.get<User[]>('/api/users', {})
    this.users = await lastValueFrom(users$)
    for(var user of this.users)
    {
      if(user.username == this.username && await compare(this.password, user.password))
      {
        this.success = true
      }
    }
    this.username = ''
    this.password = ''
    if(this.success)
    {
      this.router.navigate(['sidebar'])
    }
  }

 it('should create', () => {
    expect(component).toBeTruthy();
  });
 
 it('should create', () => {
   expect(component).toBeTruthy();
  });
  
   it("testing header", ()=>{
    const data=fixture.nativeElement;
    expect(data.querySelector(".content").textContent).toContain("This explore page holds posts/music from other artists that they do not follow")
  })

   it('should create', () => {
    expect(component).toBeTruthy();
  });
  
    it("testing header", ()=>{
    const data=fixture.nativeElement;
    expect(data.querySelector(".content").textContent).toContain("This is the home page. Users will be able to see things posted from people they follow.")
  })

   it('should create', () => {
    expect(component).toBeTruthy();
  });
  
  it("testing header", ()=>{
    const data=fixture.nativeElement;
    expect(data.querySelector(".content").textContent).toContain("Here users will be able to post/upload music")
  })

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should create the app', () => {
    const fixture = TestBed.createComponent(SideBarComponent);
    const app = fixture.componentInstance;
    expect(app).toBeTruthy();
  });

  it(should have as title 'homepage', () => {
    const fixture = TestBed.createComponent(SideBarComponent);
    const app = fixture.componentInstance;
    expect(app.title).toEqual('homepage');
  });

  it('should render title', () => {
    const fixture = TestBed.createComponent(SideBarComponent);
    fixture.detectChanges();
    const compiled = fixture.nativeElement as HTMLElement;
    expect(compiled.querySelector('.content span')?.textContent).toContain('homepage app is running!');
  });

# Backend:

    Increased security with hashed passwords

    Finished implementation for file upload and retrieval capabilities to User
    
    Set up HTTPClient for Frontend to communicate with backend
    
    Assisted Frontend with Login Page completion

    Verified that the API is working correctly by creating multiple new Go unit tests for Files and Users
    
# Backend Unit Tests:
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

    
    
# Backend Documentation
 ## User Struct
 
 **gorm.Model**:  Stores uint ID for identification, time.Time CreatedAt, time.Time UpdatedAt, and time.Time DeletedAt (initially null)<br><br>
 **username**: A string that identifies a user besides ID<br><br>
 **password**: A string to be used for login identification<br><br>
 **files**: A collection of File structs attributed to the user<br><br>
 
 ## File Struct
 
 **Filename** : A string that identifies the file     <br><br>
 **Size**     : 64-bit int that stores the file size. <br><br>
 **Type**     : A string that stores the file type.   <br><br>
 **OwnerID**  : String that corresponds to the owner. <br><br>
 **CreatedAt**: Time the file was created. 	      <br><br>
 **Data**     : Byte array that stores the data       <br><br>
 
 ## User Functions
    
 **getUser()**: Takes in a ResponseWrite and HTTP request and calls DB.First(&user, ["id"]) to select a User at ID "id". Associated with the ("/api/users/{id}", GetUser).Methods("GET") handle in initializeRouter. <br><br>
 **getUsers()**: Takes in a ResponseWrite and HTTP request and calls DB.First(&users) to select and save an array of Users. Associated with the ("/api/users", GetUsers).Methods("GET") handle in initializeRouter<br><br>
 **createUser()**: Takes in a ResponseWrite and HTTP request and calls DB.Create(&user) to post a User to the database. Associated with the ("/api/users", CreateUser).Methods("POST") handle in initializeRouter. <br><br>
 **updateUser()**: Takes in a ResponseWrite and HTTP request and calls DB.First(&user, ["id"]) to select a user of ID "id" and DB.Save(&user) to update the user. Associated with the ("/api/users/{id}", UpdateUser).Methods("PUT") handle in initializeRouter.<br><br>
 **deleteUser()**: Takes in a ResponseWrite and HTTP request and calls DB.Delete(&user, ["id"]) to delete a user of ID "id". Associated with the ("/api/users/{id}", DeleteUser).Methods("DELETE") handle in initializeRouter<br><br>
 
 ## File Functions
 
 **uploadFile()** : Takes in a ResponseWrite and HTTP request and uploads the file to a dataase. The function reads the data into a byte slice, assigns it with a "file", and uploads the copy to the database. Associated with the ("/api/users/{id}/files/upload", uploadFile).Methods("POST") handle in initializeRouter<br><br>
 **getFiles()**   : Takes in a ResponseWrite and HTTP request and retrieves all files associated with a given user account. The function sets the content type header to "application/json", extracts the owner ID from the URL parameter, queries the database for all files associated with the owner ID, and returns all the retrieved files in the response with a status code of 200. Associated with the ("/api/users/{id}/files", getFiles).Methods("GET") handle in initializeRouter<br><br>
 **updateFile()** : Takes in a ResponseWrite and HTTP request and updates a file associated with a given user account. The function sets the content type header to "application/json", decodes the request body into the file struct, extracts the owner ID and file ID from the URL parameter, queries the database for the file associated with the owner ID and the file ID, updates only the non-empty fields of the file struct in the database, and returns the updated file in the response with a status code of 200. Associated with the ("/api/users/{id}/files/{fid}", updateFile).Methods("PUT") handle in initializeRouter<br><br>
 **deleteFile()** : Takes in a ResponseWrite and HTTP request and deletes a specified file associated with a given user account. The function sets the content type header to "application/json", extracts the owner ID and file ID from the URL parameter, queries the database for the file associated with the owner ID and the file ID, deletes the file from the database, and returns a confirmation message in the response with a status code of 200. Associated with the ("/api/users/{id}/files/{fid}", deleteFile).Methods("DELETE") handle in initializeRouter<br><br>
 
 
 ## Database Initialization Functions
 
 **main()**: The entry point of the program. Calls **initialMigration** to set up database and **initializeRouter** to create handler functions.<br><br>
 **initialMigration()**: Opens a new MySQL server using GORM, panics on error. Automigrates users and files so tables are up to date.<br><br>
 **initializeRouter()**: Creates router using Gorilla MUX and desired functions. Listens on port :5000.<br><br>
