import { Component } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http'
import { lastValueFrom } from 'rxjs';
import { Router } from '@angular/router';

interface APIFile{
	filename: string
	type: string
	data: string
}

@Component({
  selector: 'app-post',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent {

  public data = ''
  public files: APIFile[] = []
  public fileName = ''
  public uploadfile: File
  public filedisplay: APIFile = {
    filename: '',
    type: '',
    data: ''
  }
  public filedisplay2: APIFile = {
    filename: '',
    type: '',
    data: ''
  }
  constructor(
    private httpClient: HttpClient,
    private router: Router
  ){}

  onFileSelected(event) {
    this.uploadfile = event.target.files[0]
}

  uploadFile()
  {
    const formData = new FormData()
    formData.append('file', this.uploadfile)
    this.httpClient.post("/api/users/1/files/upload", formData).subscribe()

  }

  
  async getFiles()
  {
    const files$ = await this.httpClient.get<APIFile[]>('/api/users/1/files', {})
    this.files = await lastValueFrom(files$)
    for(var file of this.files)
    {
      if(file.type == "image/png")
      {
        console.log("image recognized")
        this.filedisplay = file
      }
      if(file.type == "audio/mpeg")
      {
        console.log("audio recognized")
        this.filedisplay2 = file
      }
      console.log(file.filename)
    }
    
  }

}