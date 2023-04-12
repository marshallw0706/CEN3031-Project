import { Component } from '@angular/core';
import { GlobalConstants } from 'src/common/global-constants';
import { HttpClient, HttpHeaders } from '@angular/common/http'
import { lastValueFrom } from 'rxjs';
import { Router } from '@angular/router';

interface APIFile{
  ID: BigInt
	filename: string
	type: string
	data: string
}

interface User{
  ID: BigInt
  username: string
  password: string
}

interface ProfileInfo{
  name: string
  jobtitle: string
  description: string
}

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent {
  public userID = GlobalConstants.loggedinid
  public handle: string
  public files: APIFile[] = []
  public user: User
  public profileinfo: ProfileInfo
  public uploadfile: File
  public mostrecentfiletype: string
  public mostrecentfileid: BigInt
  constructor(
    private httpClient: HttpClient,
    private router: Router
  ){}

  profile = {
    name: 'Sound Space  User',
    jobTitle: 'Master Musician',
    userImage: 'https://via.placeholder.com/350x150',
    description: 'Insert profile page description here.'
  };
  editMode = false;

  ngOnInit(): void {
    this.getUsertoView()
    this.getFiles()
  }

  async getUsertoView()
  {
    const user$ = await this.httpClient.get<User>('/api/users/'+GlobalConstants.viewprofileid, {})
    this.user = await lastValueFrom(user$)
    this.handle = this.user.username
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
}