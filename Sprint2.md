# Front-End:

	Updated login/sign up pages and made sure they were connected with each other

	Added a new ‘user profile’ page wheres users can view their profile and will eventually be able to edit and share it

	A page with a navigation bar was created so that users can easily access other parts of the application

	Cypress Test - For this sprint the test focused on ensuring that each of our pages were loading in correctly.
	
	Unit Test - 


# Backend:

    Changed User values to more accurately reflect user data (from "name", "email" to "username", "password", "files")

    Added file upload and retrieval capabilities to User

    Verified that the API is working correctly by creating GET, POST, and PUT requests via Go unit tests
    
    
# Backend Documentation
 ## User Struct
 
 **gorm.Model**:  Stores uint ID for identification, time.Time CreatedAt, time.Time UpdatedAt, and time.Time DeletedAt (initially null)<br><br>
 **username**: A string that identifies a user besides ID<br><br>
 **password**: A string to be used for login identification<br><br>
 
 //add others
 
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
