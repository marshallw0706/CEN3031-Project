# Front-End:

	Updated login/sign up pages and made sure they were connected with each other. We also made sure to connect the back end
	
	with the front end for the login pages. Now when a user signs up or logins in, this information is stored in the back end. 

	Added a new ‘user profile’ page wheres users can view their profile and will eventually be able to edit and share it. There
	
	is a place for a user to upload a photo as well.

	A single page application with a navigation bar was created so that users can easily access other parts of the application. The 
	
	navigation bar will allow users to go to the home, profile, explore, and post part of the application. In this sprint, users
	
	can click between each tab and a small description of what they will hold in the future is provided. A logo was also added
	
	to the navigation bar. A component and route was declared for each of the tabs on the navigation bar.

	Cypress Test - For this sprint the test focused on ensuring that each of our pages were loading in correctly. (The test itself was
	
	shown in the video)
	
	Unit Test - Below is the unit tests that were created for the front end. They are shown in greater detail in the video. All
	
	of these tests passed with no failures. 
	
	This tests whether or not the app runs (Provided by Angular automatically)
	
	  it('should create the app', () => {
    		const fixture = TestBed.createComponent(AppComponent);
   		const app = fixture.componentInstance;
    		expect(app).toBeTruthy();
  	  });
	  
	  This tests if the side bar was created. This is done in each one of the components  (Provided by Angular automatically)
	  
	  it('should create', () => {
    		expect(component).toBeTruthy();
          );
	  
	 This tests if when the profile tab is clicked the messsage displayed shows 
	 
 	 it("testing header", ()=>{
    		const data=fixture.nativeElement;
   		expect(data.querySelector(".content").textContent).toContain("This is the users profile page")
   	  })
	  
	 This tests if when the post tab is clicked the messsage displayed shows 
	   
	 it("testing header", ()=>{
    		const data=fixture.nativeElement;
   	 	expect(data.querySelector(".content").textContent).toContain("Here users will be able to post/upload music")
	  })
  
        This tests if when the home tab is clicked the messsage displayed shows 
	
    	it("testing header", ()=>{
   		 const data=fixture.nativeElement;
    		 expect(data.querySelector(".content").textContent).toContain("This is the home page. Users will be able to see things 				         posted from people they follow.")
 	 })
  
   	This tests if when the explore tab is clicked the messsage displayed shows 
	
    	it("testing header", ()=>{
    		const data=fixture.nativeElement;
    		expect(data.querySelector(".content").textContent).toContain("This explore page holds posts/music from other artists 		     that 	          they do not follow")
 	 })

# Backend:

    Changed User values to more accurately reflect user data (from "name", "email" to "username", "password", "files")

    Added file upload and retrieval capabilities to User

    Verified that the API is working correctly by creating GET, POST, and PUT requests via Go unit tests
    
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
