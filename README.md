# Project Name: SoundSpace
## Project Description 
A social media platform, allowing for musicians to connect and share music. This social media platform will allow users to create their own profile, upload their own music, and interact with similar musicians. The users will be able to interact with the music from other creators and leave suggestions in a comment section.
## Members 
Sofia Celestrin - Front-End, <br><br>
Wendell Marshall - Front-End,<br><br>
Ryan Campisi - Back-End,<br><br>
Jonathan Santiago - Back-End

# How to Run
## Specifications
SoundSpace runs using MySQL for the backend and Angular Client for the frontend. Any device capable of running both software is capable of running SoundSpace. Reference MySQL and Angular for more information: https://dev.mysql.com/doc/refman/8.0/en/what-is-mysql.html https://angular.io/guide/what-is-angular
## Step 1 
Make sure MySQL server has been downloaded on your computer. It is recommended to install MySQL workbench. https://dev.mysql.com/downloads/installer/
<br><br>
<img src="/gifs/downloadmysql.gif" width="600" height="300"/>
## Step 2
Start a new server on host:localhost with username root and no password. MySQL workbench makes this easy as it is typically the default.<br><br>
*If you accidentally create a password and are having trouble accessing the database, try editing the password using an ALTER USER root:localhost IDENTIFIED by '(password)' query*
<br><br>
<img src="/gifs/mysqlworkbench.gif" width="600" height="300"/>
## Step 3
Navigate to the 'api' folder and run the api by building the project (using 'go build') and running the executable (using .\api.exe on Windows). Then navigate to the 'client' folder and run the client by running "npm run start".<br><br> *If you run into any errors running the API, check the error message and make sure your MySQL server is running properly. It is listening on Port 5000 which may cause problems on select devices. If you run into errors running the client, it is most likely due to uninstalled packages. Just run 'npm install' or any other specified installations listed in the error message*.
<br><br>
<img src="/gifs/runproject (1).gif" width="600" height="300"/>
## Step 4
Once both are running, navigate to 'localhost:4200' in any web browser to use SoundSpace.
<br><br>
<img src="/gifs/localhost.gif" width="600" height="300"/>

