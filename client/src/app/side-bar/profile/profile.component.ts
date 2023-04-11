import { Component } from '@angular/core';
import { GlobalConstants } from 'src/common/global-constants';
import { HttpClient, HttpHeaders } from '@angular/common/http'
import { lastValueFrom } from 'rxjs';
import { Router } from '@angular/router';

interface APIFile{
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
    this.httpClient.post("/api/users/"+GlobalConstants.loggedinid+"/files/upload", formData).subscribe()

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
      if(file.type == "image/png" || file.type == "image/jpg"  && file.filename == "profilepic.png")
      {
        console.log("image recognized")
        this.profile.userImage = "data:" + file.type + ";base64," + file.data
        console.log(this.profile.userImage)
      }
      console.log(file.filename)
    }
    
  }
}