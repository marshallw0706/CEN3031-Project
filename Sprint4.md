# Front-End

In this sprint we focused on minor issues that we had in Sprint 3 and added functionality to the app. 

We changed the layout and writing in all of the headers and added a slogan. 

We updated the mat-icon in the side navigation to better fit what they led to. 

We changed the explore page to a trending page that shows users the top 20 artists of the day. We added 20 different artists that have an actual user name and it has a picture of the album/single with a small caption. The user is also able to click the Go to Profile button and it will lead them to the profile of the user.

We also created an about us button on the sidbar that leads the user to a pafe that gives a short desciption on what the app is and how they can be featured in the top 20 artists page. 

We fixed up the way the profile page looks. We placed the profile of the user on a mat-card and centered it. Users can click on the pencil next to their name and edit the information including the picture.

The home page was also fixed so that it shows on the same page of the side bar. Each post shows up on a mat-card and the user can click on the like button to like a photo, the comment button to comment, and can click on the username to go to the profile of the user.

The post button was also fixed so that it shows up on the home page.

The user can now click the logout button to logout

User can use search bar to search for another user.

Theme colors were changed to match the theme of the whole app. 

Front-End Tests:

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


it('should create the app', () => {
    const fixture = TestBed.createComponent(SideBarComponent);
    const app = fixture.componentInstance;
    expect(app).toBeTruthy();
  });

  it(`should have as title 'homepage'`, () => {
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

async gotoprofile()
  {
    const users$ = await this.httpClient.get<User[]>('/api/users', {})
    const users = await lastValueFrom(users$)
    for(var user of users)
    {
      if(user.username == this.usernametoprofile)
      {
        GlobalConstants.viewprofileid = user.ID
        this.router.navigate(['profile'])
      }
    }
  }

    togglefollow()
  {
    this.user.following = !this.user.following
    if(this.user.following)
    {
        this.httpClient.post('/api/users/'+GlobalConstants.loggedinid+'/follow/'+GlobalConstants.viewprofileid, {}).subscribe(
          (response: any) => console.log("successful follow: " + response),
      (error) => console.log("failure to follow: " + error)
    );
    }
    if(!this.user.following)
    {
      this.httpClient.delete('/api/users/'+GlobalConstants.loggedinid+'/unfollow/'+GlobalConstants.viewprofileid, {}).subscribe(
        (response: any) => console.log("successful unfollow: " + response),
    (error) => console.log("failure to unfollow: " + error)
  );
    }
  }

    async getFiles()
  {
    const files$ = await this.httpClient.get<APIFile[]>('/api/users/'+GlobalConstants.viewprofileid+'/files', {})
    this.files = await lastValueFrom(files$)
    for(var file of this.files)
    {
      if((file.type == "image/png" || file.type == "image/jpg" || file.type == "image/jpeg")  && file.filename == "profilepic.png")
      {
        console.log("image recognized")
        this.profile.userImage = "data:" + file.type + ";base64," + file.data
        console.log(this.profile.userImage)
      }
    }
    const profileinfo$ = await this.httpClient.get<ProfileInfo>('/api/users/'+GlobalConstants.viewprofileid+'/profile', {})
    this.profileinfo = await lastValueFrom(profileinfo$)
    console.log("this is " + this.profileinfo.name)
    if(this.profileinfo.description != "")
      this.profile.description = this.profileinfo.description
    if(this.profileinfo.jobtitle != "")
      this.profile.jobTitle = this.profileinfo.jobtitle
    if(this.profileinfo.name != "")
      this.profile.name = this.profileinfo.name
    
  }

   async addLike(likefile: APIFile)

    likefile.liked = !likefile.liked;

  if (likefile.liked) {
    likefile.likes++
    this.httpClient.post('/api/users/' + GlobalConstants.loggedinid + '/like/' + likefile.owner_id + '/' + likefile.ID, {}).subscribe(
      (response: any) => console.log("successful like: " + response),
      (error) => console.log("failure to like: " + error)
    );
  } else {
    likefile.likes--
    this.httpClient.delete('/api/users/' + GlobalConstants.loggedinid + '/unlike/' + likefile.owner_id + '/' + likefile.ID, {}).subscribe(
      (response: any) => console.log("successful unlike: " + response),
      (error) => console.log("failure to unlike: " + error)
    );
  }

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

  Cypress Test:
  describe('check login page', () => {
  it('passes', () => {
    cy.visit('http://localhost:4200')
  })
})

describe('check sign up page', () => {
  it('passes', () => {
    cy.visit('http://localhost:4200/signup')
  })
})

describe('check sidebar', () => {
  it('passes', () => {
    cy.visit('http://localhost:4200/sidebar')
  })
})

describe('check profile', () => {
  it('passes', () => {
    cy.visit('http://localhost:4200/profile')
  })
})

describe('check post', () => {
  it('passes', () => {
    cy.visit('http://localhost:4200/post')
  })
})

andleHomeButtonClick() {
    if (this.isPostVisible) {
      this.togglePostVisibility();
    }
    this.homeComponent.getFiles();
  }

  async gotoprofile()
  {
    const users$ = await this.httpClient.get<User[]>('/api/users', {})
    const users = await lastValueFrom(users$)
    for(var user of users)
    {
      if(user.username == this.usernametoprofile)
      {
        GlobalConstants.viewprofileid = user.ID
        this.router.navigate(['profile'])
      }
    }
  }

  togglePostVisibility() {
    this.isPostVisible = !this.isPostVisible;
  }

describe('SearchingComponent', () => {
  let component: SearchingComponent;
  let fixture: ComponentFixture<SearchingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SearchingComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SearchingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

  async getUsertoView()
  {
    const user$ = await this.httpClient.get<User>('/api/users/'+GlobalConstants.viewprofileid, {})
    this.user = await lastValueFrom(user$)
    this.handle = this.user.username
    const users$ = this.httpClient.get<User[]>('/api/users/'+GlobalConstants.loggedinid+'/following', {})
    this.users = await lastValueFrom(users$)
    if(this.users != null)
    {
    for(var user of this.users)
    {
      console.log("Looking at users")
      if(user.ID == GlobalConstants.viewprofileid)
      {
        this.user.following = true
      }
    }
  }
    if(this.user.following == null)
    {
      this.user.following = false
    }
  }

  onFileSelected(event) {
    this.uploadfile = event.target.files[0]
}

  uploadFile()
  {
    const formData = new FormData()
    formData.append('file', this.uploadfile)
    this.httpClient.post("/api/users/"+GlobalConstants.viewprofileid+"/files/upload", formData).subscribe(
      (response: any) => {this.mostrecentfileid = response.ID; 
      this.mostrecentfiletype = response.type;
    if(this.mostrecentfiletype != "image/png" && this.mostrecentfiletype != "image/jpg" && this.mostrecentfiletype != "image/jpeg")
    {
      console.log(this.mostrecentfiletype + " not a valid type, deleting now")
      this.httpClient.delete("/api/users/"+GlobalConstants.viewprofileid+"/files/"+this.mostrecentfileid, {}).subscribe()
    }
    else{
    this.httpClient.put("/api/users/"+GlobalConstants.viewprofileid+"/files/"+this.mostrecentfileid, {"filename": "profilepic.png"}).subscribe()
    }
    this.getFiles()}, 
      (error) => console.log("bad")
    )

  }

  toggleEditMode() {
    if(this.editMode)
    {
      this.httpClient.put("/api/users/"+GlobalConstants.viewprofileid+"/profile", {
        "name": this.profile.name,
        "jobtitle": this.profile.jobTitle,
        "description": this.profile.description
    }).subscribe((response: any) => console.log("good update profileinfo => " + response + " name: " + response.name), (error) => console.log("bad update profileinfo " + error))
    }
    this.editMode = !this.editMode;
  }

  togglefollow()
  
    this.user.following = !this.user.following
    if(this.user.following)
    {
        this.httpClient.post('/api/users/'+GlobalConstants.loggedinid+'/follow/'+GlobalConstants.viewprofileid, {}).subscribe(
          (response: any) => console.log("successful follow: " + response),
      (error) => console.log("failure to follow: " + error)
    );
    }
    if(!this.user.following)
    
      this.httpClient.delete('/api/users/'+GlobalConstants.loggedinid+'/unfollow/'+GlobalConstants.viewprofileid, {}).subscribe(
        (response: any) => console.log("successful unfollow: " + response),
    (error) => console.log("failure to unfollow: " + error)
  );
    
  

# Backend:

    Added ability to save profile info

    Added ability to follow/unfollow another User
    
    Added ability to like/unlike a file
    
    Added ability to post/delete comments on a file
    
    Assisted Frontend in making backend requests on Angular

    Verified that the API is working correctly by creating multiple new Go unit tests for Files, Users, Comments, Likes, and Followers
    
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

	func TestDeleteComment(t *testing.T) {
		InitialMigration()
		router := mux.NewRouter()
		router.HandleFunc("/api/users/{uid}/comment/{id}/{fid}/{cid}", deleteComment).Methods("DELETE")

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
		req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/users/%d/comment/%d/%d/%d", user.ID, user.ID, file.ID, comment.ID), nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusOK)
		}

		var response map[string]string
		err = json.Unmarshal(rr.Body.Bytes(), &response)
		if err != nil {
			t.Fatal(err)
		}

		if response["message"] != "comment deleted" {
			t.Errorf("handler returned unexpected message: got %v want %v",
				response["message"], "comment deleted")
		}

		// Check if the comment was actually deleted
		var deletedComment Comment
		if err := DB.First(&deletedComment, comment.ID).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
			t.Errorf("handler did not delete comment as expected")
		}

		// Clean up test data
		DB.Delete(&file)
		DB.Delete(&user)
	}

	func TestPostComment(t *testing.T) {
		InitialMigration()
		router := mux.NewRouter()
		router.HandleFunc("/api/users/{uid}/comment/{id}/{fid}", postComment).Methods("POST")

		user := &User{Username: "testuser", Password: "testpass"}
		DB.Create(&user)

		file := &File{
			Filename:    "testfile",
			Size:        1024,
			Type:        "text/plain",
			OwnerID:     strconv.Itoa(int(user.ID)),
			CreatedAt:   time.Now().Unix(),
			Data:        []byte("test data"),
			Likes:       0,
			Description: "Test file",
		}
		DB.Create(&file)

		commentData := `{"content": "test comment"}`
		req, err := http.NewRequest("POST", fmt.Sprintf("/api/users/%d/comment/%d/%d", user.ID, user.ID, file.ID), strings.NewReader(commentData))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusCreated)
		}

		var returnedComment Comment
		json.Unmarshal(rr.Body.Bytes(), &returnedComment)
		if returnedComment.Content != "test comment" {
			t.Errorf("handler returned wrong comment content: got %v want %v",
				returnedComment.Content, "test comment")
		}

		DB.Delete(&returnedComment)
		DB.Delete(&file)
		DB.Delete(&user)
	}

	func TestGetLikedByUsers(t *testing.T) {
		InitialMigration()
		router := mux.NewRouter()
		router.HandleFunc("/api/users/{id}/files/{fid}/likedby", getLikedByUsers).Methods("GET")

		// Create test data
		user := &User{Username: "testuser", Password: "testpass"}
		DB.Create(user)

		liker := &User{Username: "liker", Password: "testpass"}
		DB.Create(liker)

		file := &File{
			Filename:    "testfile",
			Size:        1024,
			Type:        "text/plain",
			OwnerID:     strconv.Itoa(int(user.ID)),
			CreatedAt:   time.Now().Unix(),
			Data:        []byte("Test data"),
			Likes:       1,
			Description: "Test description",
			LikedBy:     []User{*liker},
		}
		DB.Create(file)

		// Test the endpoint
		req, err := http.NewRequest("GET", fmt.Sprintf("/api/users/%d/files/%d/likedby", user.ID, file.ID), nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusOK)
		}

		var likedByUsers []User
		err = json.Unmarshal(rr.Body.Bytes(), &likedByUsers)
		if err != nil {
			t.Fatal(err)
		}

		if len(likedByUsers) != 1 || likedByUsers[0].Username != liker.Username {
			t.Errorf("handler returned unexpected liked by users: got %v want %v",
				likedByUsers, []User{*liker})
		}

		// Clean up test data
		DB.Delete(&file)
		DB.Delete(&liker)
		DB.Delete(&user)
	}

	func TestGetFollowingUsers(t *testing.T) {
		InitialMigration()
		router := mux.NewRouter()
		router.HandleFunc("/api/users/{id}/following", getFollowingUsers).Methods("GET")

		// Create test data
		user := &User{Username: "testuser", Password: "testpass"}
		DB.Create(user)

		followed := &User{Username: "followeduser", Password: "testpass"}
		DB.Create(followed)

		user.Following = append(user.Following, *followed)
		DB.Save(user)

		// Test the endpoint
		req, err := http.NewRequest("GET", fmt.Sprintf("/api/users/%d/following", user.ID), nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusOK)
		}

		var followingUsers []User
		err = json.Unmarshal(rr.Body.Bytes(), &followingUsers)
		if err != nil {
			t.Fatal(err)
		}

		if len(followingUsers) != 1 || followingUsers[0].Username != followed.Username {
			t.Errorf("handler returned unexpected following users: got %v want %v",
				followingUsers, []User{*followed})
		}

		// Clean up test data
		user.Following = []User{}
		DB.Save(user)
		DB.Delete(&followed)
		DB.Delete(&user)
	}

	func TestLikeFile(t *testing.T) {
		InitialMigration()
		router := mux.NewRouter()
		router.HandleFunc("/api/users/{uid}/like/{id}/{fid}", likeFile).Methods("POST")

		user := &User{Username: "testuser", Password: "testpass"}
		DB.Create(&user)

		file := &File{
			Filename:    "testfile",
			Size:        1024,
			Type:        "text/plain",
			OwnerID:     strconv.Itoa(int(user.ID)),
			CreatedAt:   time.Now().Unix(),
			Data:        []byte("test data"),
			Likes:       0,
			Description: "Test file",
		}
		DB.Create(&file)

		req, err := http.NewRequest("POST", fmt.Sprintf("/api/users/%d/like/%d/%d", user.ID, user.ID, file.ID), nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusOK)
		}

		var returnedFile File
		json.Unmarshal(rr.Body.Bytes(), &returnedFile)
		if returnedFile.Likes != 1 {
			t.Errorf("handler returned wrong number of likes: got %v want %v",
				returnedFile.Likes, 1)
		}

		DB.Delete(&file)
		DB.Delete(&user)
	}

	func TestUnlikeFile(t *testing.T) {
		InitialMigration()
		router := mux.NewRouter()
		router.HandleFunc("/api/users/{uid}/unlike/{id}/{fid}", unlikeFile).Methods("DELETE")

		user := &User{Username: "testuser", Password: "testpass"}
		DB.Create(&user)

		file := &File{
			Filename:    "testfile",
			Size:        1024,
			Type:        "text/plain",
			OwnerID:     strconv.Itoa(int(user.ID)),
			CreatedAt:   time.Now().Unix(),
			Data:        []byte("test data"),
			Likes:       1,
			Description: "Test file",
		}
		DB.Create(&file)

		// Add user as a liker of the file
		DB.Model(&file).Association("LikedBy").Append(user)

		req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/users/%d/unlike/%d/%d", user.ID, user.ID, file.ID), nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusOK)
		}

		var returnedFile File
		json.Unmarshal(rr.Body.Bytes(), &returnedFile)
		if returnedFile.Likes != 0 {
			t.Errorf("handler returned wrong number of likes: got %v want %v",
				returnedFile.Likes, 0)
		}

		DB.Delete(&file)
		DB.Delete(&user)
	}

	func TestRemoveFollower(t *testing.T) {
		InitialMigration()
		router := mux.NewRouter()
		router.HandleFunc("/api/users/{id}/unfollow/{fid}", removeFollower).Methods("DELETE")

		user := &User{Username: "testuser", Password: "testpass"}
		follower := &User{Username: "testfollower", Password: "testpass"}
		DB.Create(&user)
		DB.Create(&follower)

		// Make the user follow the follower before testing
		DB.Model(&user).Association("Following").Append(&follower)

		req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/users/%d/unfollow/%d", user.ID, follower.ID), nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusOK)
		}

		var returnedUser User
		json.Unmarshal(rr.Body.Bytes(), &returnedUser)
		if len(returnedUser.Following) != 0 {
			t.Errorf("handler returned wrong number of followers: got %v want %v",
				len(returnedUser.Following), 0)
		}

		DB.Delete(&follower)
		DB.Delete(&user)
	}

	func TestAddFollower(t *testing.T) {
		InitialMigration()
		router := mux.NewRouter()
		router.HandleFunc("/api/users/{id}/follow/{fid}", addFollower).Methods("POST")

		user := &User{Username: "testuser", Password: "testpass"}
		follower := &User{Username: "testfollower", Password: "testpass"}
		DB.Create(&user)
		DB.Create(&follower)

		req, err := http.NewRequest("POST", fmt.Sprintf("/api/users/%d/follow/%d", user.ID, follower.ID), nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusOK)
		}

		var returnedUser User
		json.Unmarshal(rr.Body.Bytes(), &returnedUser)
		if len(returnedUser.Following) != 1 {
			t.Errorf("handler returned wrong number of followers: got %v want %v",
				len(returnedUser.Following), 1)
		}

		// Clean up following relationship
		DB.Model(&user).Association("Following").Delete(&follower)
		DB.Delete(&follower)
		DB.Delete(&user)
	}

    
    
# Backend Documentation
 ## User Struct
 
 **gorm.Model**:  Stores uint ID for identification, time.Time CreatedAt, time.Time UpdatedAt, and time.Time DeletedAt (initially null)<br><br>
 **username**: A string that identifies a user besides ID<br><br>
 **password**: A string to be used for login identification<br><br>
 **files**: A collection of File structs attributed to the user<br><br>
 **profileinfo**: A struct representing the user's profile information <br><br>
 **following**: An array of Users that the User is following
 
 ## File Struct
 **Filename** : A string that identifies the file     <br><br>
 **Size**     : 64-bit int that stores the file size. <br><br>
 **Type**     : A string that stores the file type.   <br><br>
 **OwnerID**  : String that corresponds to the owner. <br><br>
 **CreatedAt**: Time the file was created. 	      <br><br>
 **Data**     : Byte array that stores the data       <br><br>
 **Likes**    : An integer that represents the number of likes <br><br>
 **LikedBy**  : A list of users that liked the file <br><br>
 **Description** : A string with information on the file<br><br>
 **Comments** : An array of Comment structs representing comments by users <br><br>
 
 ## Comment Struct
**gorm.Model**:  Stores uint ID for identification, time.Time CreatedAt, time.Time UpdatedAt, and time.Time DeletedAt (initially null)<br><br>
**Content**: A string representing what the commenter said<br><br>
**PostedBy**: A user representing the user who commented<br><br>
**UserID**: An integer representing the ID of the owner of the file receiving the comment<br><br>
**FileID**: And integer representing the ID of the file receiveing the comment

## ProfileInfo Struct
**OwnerID**: An integer that is copied from the User's ID field<br><br>
**Name**: A string representing the name added to the profile by the user<br><br>
**JobTitle**: A string representing the job title added to the profile by the user<br><br>
**Description**: A string representing the profile description added to the profile by the user
 
 ## User Functions
    
 **getUser()**: Takes in a ResponseWrite and HTTP request and calls DB.First(&user, ["id"]) to select a User at ID "id". Associated with the ("/api/users/{id}", GetUser).Methods("GET") handle in initializeRouter. <br><br>
 **getUsers()**: Takes in a ResponseWrite and HTTP request and calls DB.First(&users) to select and save an array of Users. Associated with the ("/api/users", GetUsers).Methods("GET") handle in initializeRouter<br><br>
 **createUser()**: Takes in a ResponseWrite and HTTP request and calls DB.Create(&user) to post a User to the database. Associated with the ("/api/users", CreateUser).Methods("POST") handle in initializeRouter. <br><br>
 **updateUser()**: Takes in a ResponseWrite and HTTP request and calls DB.First(&user, ["id"]) to select a user of ID "id" and DB.Save(&user) to update the user. Associated with the ("/api/users/{id}", UpdateUser).Methods("PUT") handle in initializeRouter.<br><br>
 **deleteUser()**: Takes in a ResponseWrite and HTTP request and calls DB.Delete(&user, ["id"]) to delete a user of ID "id". Associated with the ("/api/users/{id}", DeleteUser).Methods("DELETE") handle in initializeRouter<br><br>
 
 ## Comment Functions
    
 **postComment()**: Takes in a ResponseWrite and HTTP request and calls DB.Create(&newComment) to add a comment to a file. Associated with the ("/api/users/{uid}/comment/{id}/{fid}", postComment).Methods("POST") handle in initializeRouter. <br><br>
 **deleteComment()**: Takes in a ResponseWrite and HTTP request and calls DB.Delete(&comment) to delete a comment from a file. ("/api/users/{uid}/comment/{id}/{fid}/{cid}", deleteComment) handle in initializeRouter<br><br>
 **getComments()**: Takes in a ResponseWrite and HTTP request and calls DB.Preload("Comments.PostedBy").Where("owner_id = ? AND id = ?", owner.ID, params["fid"]).First(&file) to preload the user who posted the comment(s), and retrieve said comment(s). Associated with the ("/api/users/{id}/files/{fid}/comments", getComments).Methods("GET") handle in initializeRouter. <br><br>
 
 ## Follower Functions
  **addFollower()**: This function takes in a ResponseWriter and an HTTP request as input parameters and calls DB.Model(&user).Association("Following").Append(&follower) to add a follower to a user. It is associated with the ("/api/users/{id}/followers/{fid}", addFollower).Methods("POST") handle in initializeRouter. <br><br>
 **removeFollower()**: This function takes in a ResponseWriter and an HTTP request as input parameters and calls DB.Model(&user).Association("Following").Delete(&follower) to delete a follower from a user. It is associated with the ("/api/users/{id}/followers/{fid}", removeFollower).Methods("DELETE") handle in initializeRouter. <br><br>
 **getFollowingUsers()**:  This function takes in a ResponseWriter and an HTTP request as input parameters and calls DB.Preload("Following").First(&user, params["id"]) to get following users. It is associated with the ("/api/users/{id}/following", getFollowingUsers).Methods("GET") handle in initializeRouter. <br><br>
 
 ## Like Functions
  **likeFile()**: This function takes in a ResponseWriter and an HTTP request as input parameters and calls DB.Model(&file).Association("LikedBy").Append(&user) to add a user to the liked-by array of a file. It also increments the file's like count. It is associated with the ("/api/users/{uid}/files/{id}/{fid}/like", likeFile).Methods("POST") handle in initializeRouter. <br><br>
 **unlikeFile()**: This function takes in a ResponseWriter and an HTTP request as input parameters and calls DB.Model(&file).Association("LikedBy").Delete(&user) to delete a user from the liked-by array of a file. It also decrements the file's like count. It is associated with the ("/api/users/{uid}/files/{id}/{fid}/unlike", unlikeFile).Methods("POST") handle in initializeRouter. <br><br>
 **getLikedByUsers()**: This function takes in a ResponseWriter and an HTTP request as input parameters. It is associated with the ("/api/users/{uid}/files/{id}/{fid}/unlike", unlikeFile).Methods("POST") handle in initializeRouter. <br><br>
 
 ## Profile Info Functions
   **updateProfileInfo()**: This function takes in a ResponseWriter and an HTTP request as input parameters and calls DB.Model(&profile).Where("owner_id = ?", profile.OwnerID).Updates(profile) to update a user's profile info. It is associated with the ("/api/users/{id}/profile", UpdateProfileInfo) handle in initializeRouter. <br><br>
 **getProfileInfo()**: This function takes in a ResponseWriter and an HTTP request as input parameters and calls DB.Where("owner_id = ?", params["id"]).First(&profile) to retrieve a user's profile info . It is associated with the ("/api/users/{id}/profile", GetProfileInfo).Methods("GET") handle in initializeRouter. <br><br>
 
 ## File Functions
 
 **uploadFile()** : Takes in a ResponseWrite and HTTP request and uploads the file to a dataase. The function reads the data into a byte slice, assigns it with a "file", and uploads the copy to the database. Associated with the ("/api/users/{id}/files/upload", uploadFile).Methods("POST") handle in initializeRouter<br><br>
 **getFiles()**   : Takes in a ResponseWrite and HTTP request and retrieves all files associated with a given user account. The function sets the content type header to "application/json", extracts the owner ID from the URL parameter, queries the database for all files associated with the owner ID, and returns all the retrieved files in the response with a status code of 200. Associated with the ("/api/users/{id}/files", getFiles).Methods("GET") handle in initializeRouter<br><br>
 **updateFile()** : Takes in a ResponseWrite and HTTP request and updates a file associated with a given user account. The function sets the content type header to "application/json", decodes the request body into the file struct, extracts the owner ID and file ID from the URL parameter, queries the database for the file associated with the owner ID and the file ID, updates only the non-empty fields of the file struct in the database, and returns the updated file in the response with a status code of 200. Associated with the ("/api/users/{id}/files/{fid}", updateFile).Methods("PUT") handle in initializeRouter<br><br>
 **deleteFile()** : Takes in a ResponseWrite and HTTP request and deletes a specified file associated with a given user account. The function sets the content type header to "application/json", extracts the owner ID and file ID from the URL parameter, queries the database for the file associated with the owner ID and the file ID, deletes the file from the database, and returns a confirmation message in the response with a status code of 200. Associated with the ("/api/users/{id}/files/{fid}", deleteFile).Methods("DELETE") handle in initializeRouter<br><br>
 
 
 ## Database Initialization Functions
 
 **main()**: The entry point of the program. Calls **initialMigration** to set up database and **initializeRouter** to create handler functions.<br><br>
 **initialMigration()**: Opens a new MySQL server using GORM, panics on error. Automigrates users and files so tables are up to date.<br><br>
 **initializeRouter()**: Creates router using Gorilla MUX and desired functions. Listens on port :5000.<br><br>
