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

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent {
  public handle = GlobalConstants.loggedinuser
  public files: APIFile[] = []
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
    this.getFiles()
  }

  onFileSelected(event) {
    this.uploadfile = event.target.files[0]
}

  uploadFile()
  {
    const formData = new FormData()
    formData.append('file', this.uploadfile)
    this.httpClient.post("/api/users/"+GlobalConstants.loggedinid+"/files/upload", formData).subscribe(
      (response: any) => {this.mostrecentfileid = response.ID; 
      this.mostrecentfiletype = response.type;
    if(this.mostrecentfiletype != "image/png" && this.mostrecentfiletype != "image/jpg" && this.mostrecentfiletype != "image/jpeg")
    {
      console.log(this.mostrecentfiletype + " not a valid type, deleting now")
      this.httpClient.delete("/api/users/"+GlobalConstants.loggedinid+"/files/"+this.mostrecentfileid, {}).subscribe()
    }
    else{
    this.httpClient.put("/api/users/"+GlobalConstants.loggedinid+"/files/"+this.mostrecentfileid, {"filename": "profilepic.png"}).subscribe()
    }
    this.getFiles()}, 
      (error) => console.log("bad")
    )

  }

  toggleEditMode() {
    this.editMode = !this.editMode;
  }

  async getFiles()
  {
    const files$ = await this.httpClient.get<APIFile[]>('/api/users/'+GlobalConstants.loggedinid+'/files', {})
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
    
  }
}